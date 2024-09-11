package main

import (
	"testing"

	"github.com/samber/lo"

	"github.com/caiknife/ncmdl/entity"
)

func TestDownloadAlbum(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}

		err = DownloadAlbum(id, Path("./tmp/"))
		if err != nil {
			t.Error(err)
			return
		}
	})
}

func TestAlbumDetail_Random_Order(t *testing.T) {
	// const album = "https://music.163.com/album?id=153840942&userid=757014"
	const album = "https://music.163.com/album?id=195093556&userid=757014"
	id, err := AlbumLinkID(album)
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

	songIDs := lo.Map[*entity.Single, int](detail, func(item *entity.Single, index int) int {
		return item.ID
	})
	t.Log(songIDs)
	info, err := DownloadInfo(songIDs)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
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
