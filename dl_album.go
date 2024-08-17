package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/samber/lo"

	"github.com/caiknife/ncmdl/entity"
)

const (
	ErrAlbumIsEmpty  types.Error = "专辑歌曲为空"
	ErrAlbumDownload types.Error = "解析专辑歌曲下载链接异常"
)

func DownloadAlbum(albumID int, destDir string) (err error) {
	detail, err := AlbumDetail(albumID)
	if err != nil {
		return err
	}
	songIDs := lo.Map[*entity.Single, int](detail, func(item *entity.Single, index int) int {
		return item.ID
	})
	info, err := downloadInfo(songIDs)
	if err != nil {
		return err
	}
	if info.Len() != detail.Len() {
		return ErrAlbumDownload
	}

	err = asyncDownload(info, detail, destDir)
	if err != nil {
		return err
	}
	return nil
}

func AlbumDetail(albumID int) (r types.Slice[*entity.Single], err error) {
	logger.ConsoleLogger.Infoln("正在解析专辑，ID:", albumID)
	detail, err := api.GetAlbumDetail(*reqData, albumID)
	if err != nil {
		return nil, err
	}
	d := &entity.AlbumResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		return nil, err
	}
	if d.Songs.IsEmpty() {
		return nil, ErrAlbumIsEmpty
	}
	return d.Songs, nil
}
