package entity

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type DownloadResult struct {
	Data types.Slice[*Download] `json:"data"`
}

func (d *DownloadResult) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}
