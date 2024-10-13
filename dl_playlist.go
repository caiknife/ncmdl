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

func DownloadPlaylist(playlistID int, destDir string) (err error) {
	detail, err := PlaylistDetail(playlistID)
	if err != nil {
		err = errors.WithMessage(err, "playlist detail")
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
		return ErrPlaylistDownload
	}

	MergeURL(info, &detail)

	err = AsyncDownload(detail, destDir)
	if err != nil {
		err = errors.WithMessage(err, "async download")
		return err
	}
	return nil
}

const (
	ErrPlaylistIsEmpty  types.Error = "歌单列表为空"
	ErrPlaylistDownload types.Error = "解析歌单歌曲下载链接异常"
)

func PlaylistDetail(playlistID int) (r types.Slice[*entity.Single], err error) {
	logger.ConsoleLogger.Infoln("正在解析歌单，ID:", playlistID)
	detail, err := api.GetPlaylistDetail(*GetRequestData(), playlistID)
	if err != nil {
		err = errors.WithMessage(err, "api get playlist detail")
		return nil, err
	}
	d := &entity.PlaylistResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		err = errors.WithMessage(err, "json unmarshal")
		return nil, err
	}
	if d.Playlist.Tracks.IsEmpty() {
		return nil, ErrPlaylistIsEmpty
	}
	return d.Playlist.Tracks, nil
}
