package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type SingleResult struct {
	Songs types.Slice[*Single] `json:"songs"`
}

func (s *SingleResult) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}
