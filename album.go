package main

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
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
