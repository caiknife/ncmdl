package main

import (
	"sync"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/panjf2000/ants/v2"
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

func downloadInfo(songIDs []int) (r types.Slice[*entity.Download], err error) {
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

func asyncDownload(downloads types.Slice[*entity.Download], songs types.Slice[*entity.Single], destDir string) error {
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
			err := downloadFile(downloads[i].URL, single, destDir)
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
