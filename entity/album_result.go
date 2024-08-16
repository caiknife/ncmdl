package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type AlbumResult struct {
	Songs types.Slice[*Single] `json:"songs"`
	Album *Album               `json:"album"`
}

func (a *AlbumResult) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}
