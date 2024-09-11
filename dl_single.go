package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"

	"github.com/caiknife/ncmdl/entity"
)

const (
	ErrSongDetailIsEmpty types.Error = "单曲详情为空"
	ErrSongDownload      types.Error = "解析歌单歌曲下载链接异常"
)

func DownloadSingle(singleID int, destDir string) (err error) {
	detail, err := SingleDetail(singleID)
	if err != nil {
		err = errors.WithMessage(err, "single detail")
		return err
	}
	songIDs := []int{detail[0].ID}
	info, err := DownloadInfo(songIDs)
	if err != nil {
		err = errors.WithMessage(err, "download info")
		return err
	}
	if info.Len() != detail.Len() {
		return ErrSongDownload
	}

	MergeURL(info, &detail)

	err = AsyncDownload(detail, destDir)
	if err != nil {
		err = errors.WithMessage(err, "async download")
		return err
	}

	return nil
}

func SingleDetail(singleID int) (result types.Slice[*entity.Single], err error) {
	logger.ConsoleLogger.Infoln("正在解析单曲，ID:", singleID)
	detail, err := api.GetSongDetail(*GetRequestData(), []int{singleID})
	if err != nil {
		err = errors.WithMessage(err, "api get song detail")
		return nil, err
	}
	d := &entity.SingleResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		err = errors.WithMessage(err, "json unmarshal")
		return nil, err
	}
	if d.Songs.IsEmpty() {
		return nil, ErrSongDetailIsEmpty
	}
	return d.Songs, nil
}
