package main

import (
	"io"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/netutil"

	"github.com/caiknife/ncmdl/entity"
)

const (
	ErrSongDetailIsEmpty types.Error = "单曲详情为空"
	ErrSongDownload      types.Error = "解析歌单歌曲下载链接异常"
)

func writeTag(filePath string, single *entity.Single) error {
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

func downloadFile(url string, single *entity.Single, destDir string) error {
	path := destDir + single.SavePath()
	if !fileutil.IsExist(path) {
		err := fileutil.CreateDir("./" + path)
		if err != nil {
			return err
		}
	}
	mp3file := destDir + single.SaveFileName()
	if fileutil.IsExist(mp3file) {
		logger.ConsoleLogger.Warnln(single.FileName(), "文件已经存在")
		return nil
	}
	logger.ConsoleLogger.Infoln("开始下载文件", single.FileName())
	err := netutil.DownloadFile(mp3file, url)
	if err != nil {
		return err
	}
	err = writeTag(mp3file, single)
	if err != nil {
		return err
	}
	return nil
}

func DownloadSingle(singleID int, destDir string) (err error) {
	detail, err := SingleDetail(singleID)
	if err != nil {
		return err
	}
	songIDs := []int{detail[0].ID}
	info, err := downloadInfo(songIDs)
	if err != nil {
		return err
	}
	if info.IsEmpty() {
		return ErrSongDownload
	}

	err = asyncDownload(info, detail, destDir)
	if err != nil {
		return err
	}

	return nil
}

func SingleDetail(singleID int) (result types.Slice[*entity.Single], err error) {
	logger.ConsoleLogger.Infoln("正在解析单曲，ID:", singleID)
	detail, err := api.GetSongDetail(*GetRequestData(), []int{singleID})
	if err != nil {
		return nil, err
	}
	d := &entity.SingleResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		return nil, err
	}
	if d.Songs.IsEmpty() {
		return nil, ErrSongDetailIsEmpty
	}
	return d.Songs, nil
}
