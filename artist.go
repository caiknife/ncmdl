package ncmdl

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type ArtistInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (a *ArtistInfo) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}
