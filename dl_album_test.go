package main

import (
	"testing"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
)

func Test_album_detail(t *testing.T) {
	detail, err := api.GetAlbumDetail(*reqData, 123837273)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(detail.RawJson)
}
