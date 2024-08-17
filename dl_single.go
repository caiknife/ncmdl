package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"

	"github.com/caiknife/ncmdl/entity"
)

const (
	ErrSongDetailIsEmpty types.Error = "单曲详情为空"
	ErrSongDownload      types.Error = "解析歌单歌曲下载链接异常"
)

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
