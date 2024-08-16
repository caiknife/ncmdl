package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Download struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (d *Download) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}
