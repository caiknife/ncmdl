package main

import (
	"testing"
)

func TestDownloadAlbum(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}

		err = DownloadAlbum(id, "./tmp/")
		if err != nil {
			t.Error(err)
			return
		}
	})
}

func TestAlbumDetail(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := AlbumDetail(id)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}
