package main

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
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
