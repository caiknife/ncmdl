package main

import (
	"testing"
)

func TestPlaylistDetail(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		id, err := PlaylistLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := PlaylistDetail(id)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}

func TestDownloadPlaylist(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		id, err := PlaylistLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		err = DownloadPlaylist(id, "./tmp/")
	})
}
