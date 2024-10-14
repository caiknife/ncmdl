package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
)

type AlbumInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`
}

func (a *AlbumInfo) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

type AlbumResult struct {
	Songs types.Slice[*SingleInfo] `json:"songs"`
	Album *AlbumInfo               `json:"album"`
}

func (a *AlbumResult) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

func AlbumDetail(albumID int, reqData *utils.RequestData) (result types.Slice[*SingleInfo], err error) {
	detail, err := api.GetAlbumDetail(*reqData, albumID)
	if err != nil {
		err = errors.WithMessage(err, "api get album detail")
		return nil, err
	}

	d := &AlbumResult{}
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
