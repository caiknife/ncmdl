package main

import (
	"io"
	"path/filepath"
	"sync"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type DownloadInfo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (d *DownloadInfo) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}

type DownloadResult struct {
	Data types.Slice[*DownloadInfo] `json:"data"`
}

func (d *DownloadResult) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}

func Path(destDir string) string {
	abs, err := filepath.Abs(destDir)
	if err != nil {
		return "."
	}
	return abs
}

func MergeURL(downloadInfo types.Slice[*DownloadInfo], detail *types.Slice[*SingleInfo]) {
	infoMap := lo.SliceToMap[*DownloadInfo, int, *DownloadInfo](downloadInfo, func(item *DownloadInfo) (int, *DownloadInfo) {
		return item.ID, item
	})

	detail.ForEach(func(single *SingleInfo, i int) {
		if info, ok := infoMap[single.ID]; ok {
			single.URL = info.URL
		}
	})
}

func DownloadLink(songIDs []int, reqData *utils.RequestData) (result types.Slice[*DownloadInfo], err error) {
	url, err := api.GetSongURL(*reqData, api.SongURLConfig{
		EncodeType: "mp3",
		Level:      "higher",
		Ids:        songIDs,
	})
	if err != nil {
		err = errors.WithMessage(err, "api get song url")
		return nil, err
	}

	d := &DownloadResult{}
	err = fjson.UnmarshalFromString(url.RawJson, d)
	if err != nil {
		err = errors.WithMessage(err, "json unmarshal")
		return nil, err
	}
	return d.Data, nil
}

func AsyncDownload(songs types.Slice[*SingleInfo], destDir string) error {
	pool, err := ants.NewPool(defaultPoolSize)
	if err != nil {
		err = errors.WithMessage(err, "ants pool init")
		return err
	}
	defer pool.Release()

	wg := &sync.WaitGroup{}
	songs.ForEach(func(info *SingleInfo, i int) {
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			err := DownloadFile(info.URL, info, destDir)
			if err != nil {
				err = errors.WithMessage(err, "download file")
				appLogger.Errorln(err)
				return
			}
		})
		if err != nil {
			err = errors.WithMessage(err, "ant pool submit task")
			appLogger.Errorln(err)
		}
	})
	wg.Wait()
	return nil
}

func DownloadFile(url string, single *SingleInfo, destDir string) error {
	if url == "" {
		return errors.New("url is empty")
	}

	path := filepath.Join(destDir, single.SavePath())
	if !fileutil.IsExist(path) {
		err := fileutil.CreateDir(path)
		if err != nil {
			err = errors.WithMessage(err, "create dir")
			return err
		}
	}

	mp3File := filepath.Join(destDir, single.SaveFileName())
	if fileutil.IsExist(mp3File) {
		appLogger.Warnln(single.FileName(), "文件已经存在")
		return nil
	}

	appLogger.Infoln("开始下载文件", single.FileName())
	err := netutil.DownloadFile(mp3File, url)
	if err != nil {
		err = errors.WithMessage(err, "net download music file")
		return err
	}

	err = WriteTag(mp3File, single)
	if err != nil {
		err = errors.WithMessage(err, "write tag")
		return err
	}
	return nil
}

func WriteTag(filePath string, single *SingleInfo) error {
	appLogger.Infoln("正在写入标签", single.FileName())
	open, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		err = errors.WithMessage(err, "id3 open file")
		return err
	}
	defer open.Close()

	open.SetDefaultEncoding(id3v2.EncodingUTF8)
	open.SetAlbum(single.Album.Name)
	open.SetTitle(single.Name)
	open.SetArtist(single.AllArtistsTag())

	// 专辑封面图片
	response, err := netutil.HttpGet(single.Album.PicURL)
	if err != nil {
		err = errors.WithMessage(err, "net http get album pic url")
		return err
	}
	defer response.Body.Close()

	pic, err := io.ReadAll(response.Body)
	if err != nil {
		err = errors.WithMessage(err, "read album pic content")
		return err
	}
	cover := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: single.Album.Name,
		Picture:     pic,
	}
	open.AddAttachedPicture(cover)

	err = open.Save()
	if err != nil {
		err = errors.WithMessage(err, "save id3")
		return err
	}
	return nil
}
