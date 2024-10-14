package main

import (
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type DownloadInfo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (d *DownloadInfo) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}

type DownloadResult struct {
	Data types.Slice[*DownloadInfo] `json:"data"`
}

func (d *DownloadResult) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}

func Path(destDir string) string {
	abs, err := filepath.Abs(destDir)
	if err != nil {
		return "."
	}
	return abs
}
