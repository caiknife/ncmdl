package main

import (
	"io"
	"path/filepath"
	"sync"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/panjf2000/ants/v2"

	"github.com/caiknife/ncmdl/entity"
)

func DownloadInfo(songIDs []int) (r types.Slice[*entity.Download], err error) {
	url, err := api.GetSongURL(*GetRequestData(), api.SongURLConfig{
		EncodeType: "",
		Level:      "higher",
		Ids:        songIDs,
	})
	if err != nil {
		return nil, err
	}
	d := &entity.DownloadResult{}
	err = fjson.UnmarshalFromString(url.RawJson, d)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}

func AsyncDownload(downloads types.Slice[*entity.Download], songs types.Slice[*entity.Single], destDir string) error {
	pool, err := ants.NewPool(defaultPoolSize)
	if err != nil {
		return err
	}
	defer pool.Release()
	wg := &sync.WaitGroup{}
	songs.ForEach(func(single *entity.Single, i int) {
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			err := DownloadFile(downloads[i].URL, single, destDir)
			if err != nil {
				logger.ConsoleLogger.Errorln(err)
				return
			}
		})
		if err != nil {
			logger.ConsoleLogger.Errorln(err)
		}
	})
	wg.Wait()
	return nil
}

func WriteTag(filePath string, single *entity.Single) error {
	logger.ConsoleLogger.Infoln("正在写入标签", single.FileName())
	open, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer open.Close()

	open.SetAlbum(single.Album.Name)
	open.SetTitle(single.Name)
	open.SetArtist(single.Artist[0].Name)

	response, err := netutil.HttpGet(single.Album.PicURL)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	cover := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: single.Album.Name,
		Picture:     all,
	}
	open.AddAttachedPicture(cover)
	err = open.Save()
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(url string, single *entity.Single, destDir string) error {
	path := filepath.Join(destDir, single.SavePath())
	if !fileutil.IsExist(path) {
		err := fileutil.CreateDir(path)
		if err != nil {
			return err
		}
	}
	mp3file := filepath.Join(destDir, single.SaveFileName())
	if fileutil.IsExist(mp3file) {
		logger.ConsoleLogger.Warnln(single.FileName(), "文件已经存在")
		return nil
	}
	logger.ConsoleLogger.Infoln("开始下载文件", single.FileName())
	err := netutil.DownloadFile(mp3file, url)
	if err != nil {
		return err
	}
	err = WriteTag(mp3file, single)
	if err != nil {
		return err
	}
	return nil
}

func Path(destDir string) string {
	abs, err := filepath.Abs(destDir)
	if err != nil {
		return "."
	}
	return abs
}
