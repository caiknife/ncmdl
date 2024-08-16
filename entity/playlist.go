package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Playlist struct {
	Tracks types.Slice[*Single] `json:"tracks"`
}

func (p *Playlist) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}
