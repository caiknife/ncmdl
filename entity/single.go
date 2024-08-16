package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Single struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Artist types.Slice[*Artist] `json:"ar"`
	Album  *Album               `json:"al"`
}

func (s *Single) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}
