package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (a *Artist) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}
