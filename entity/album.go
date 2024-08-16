package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Album struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`
}

func (a *Album) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}
