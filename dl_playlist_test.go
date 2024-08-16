package main

import (
	"testing"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
)

func Test_playlist_detail(t *testing.T) {
	detail, err := api.GetPlaylistDetail(*reqData, 6623715587)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(detail.RawJson)
}
