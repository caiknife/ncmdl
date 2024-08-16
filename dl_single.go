package main

import (
	"io"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/netutil"

	"github.com/caiknife/ncmdl/entity"
)

const (
	br128 = 128000
	br192 = 192000
	br256 = 256000
	br320 = 320000
	br999 = 999000
)
const (
	ErrSongDetailIsEmpty   types.Error = "单曲详情为空"
	ErrSongDownloadIsEmpty types.Error = "单曲下载链接为空"
)

func writeTag(filePath string, single *entity.Single) error {
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
	err := netutil.DownloadFile(destDir+single.SaveFileName(), url)
	if err != nil {
		return err
	}
	err = writeTag(destDir+single.SaveFileName(), single)
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
	url, err := api.GetSongURL(*GetRequestData(), api.SongURLConfig{
		EncodeType: "",
		Level:      "higher",
		Ids:        songIDs,
	})
	if err != nil {
		return err
	}
	d := &entity.DownloadResult{}
	err = fjson.UnmarshalFromString(url.RawJson, d)
	if err != nil {
		return err
	}
	if d.Data.IsEmpty() {
		return ErrSongDownloadIsEmpty
	}
	e := d.Data[0].URL

	err = downloadFile(e, detail[0], destDir)
	if err != nil {
		return err
	}

	return nil
}

func SingleDetail(singleID int) (result types.Slice[*entity.Single], err error) {
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
