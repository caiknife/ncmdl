package ncmdl

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
)

type PlaylistInfo struct {
	Tracks types.Slice[*SingleInfo] `json:"tracks"`
}

func (p *PlaylistInfo) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}

type PlaylistResult struct {
	Playlist *PlaylistInfo `json:"playlist"`
}

func (p *PlaylistResult) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}

func PlaylistDetail(playlistID int, reqData *utils.RequestData) (result types.Slice[*SingleInfo], err error) {
	detail, err := api.GetPlaylistDetail(*reqData, playlistID)
	if err != nil {
		err = errors.WithMessage(err, "api get playlist detail")
		return nil, err
	}

	d := &PlaylistResult{}
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
