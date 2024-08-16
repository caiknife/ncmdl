package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type PlaylistResult struct {
	Playlist *Playlist `json:"playlist"`
}

func (p *PlaylistResult) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}
