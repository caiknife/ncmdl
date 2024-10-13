package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/caiknife/ncmdl/v2/entity"
)

const (
	ErrAlbumIsEmpty  types.Error = "专辑歌曲为空"
	ErrAlbumDownload types.Error = "解析专辑歌曲下载链接异常"
)

func DownloadAlbum(albumID int, destDir string) (err error) {
	detail, err := AlbumDetail(albumID)
	if err != nil {
		err = errors.WithMessage(err, "album detail")
		return err
	}
	songIDs := lo.Map[*entity.Single, int](detail, func(item *entity.Single, index int) int {
		return item.ID
	})
	info, err := DownloadInfo(songIDs)
	if err != nil {
		err = errors.WithMessage(err, "download info")
		return err
	}
	if info.Len() != detail.Len() {
		return ErrAlbumDownload
	}

	MergeURL(info, &detail)

	err = AsyncDownload(detail, destDir)
	if err != nil {
		err = errors.WithMessage(err, "async download")
		return err
	}
	return nil
}

func AlbumDetail(albumID int) (r types.Slice[*entity.Single], err error) {
	logger.ConsoleLogger.Infoln("正在解析专辑，ID:", albumID)
	detail, err := api.GetAlbumDetail(*GetRequestData(), albumID)
	if err != nil {
		err = errors.WithMessage(err, "api get album detail")
		return nil, err
	}
	d := &entity.AlbumResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		err = errors.WithMessage(err, "json unmarshal")
		return nil, err
	}
	if d.Songs.IsEmpty() {
		return nil, ErrAlbumIsEmpty
	}
	return d.Songs, nil
}
