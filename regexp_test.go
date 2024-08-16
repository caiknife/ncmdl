package main

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/stretchr/testify/assert"
)

var (
	albumURLs = types.Slice[string]{
		"https://music.163.com/#/album?id=123837273",
		"https://music.163.com/album?id=195093556&userid=757014",
	}

	playlistURLs = types.Slice[string]{
		"https://music.163.com/#/playlist?id=6623715587",
		"https://music.163.com/playlist?id=7350047928&userid=757014",
	}

	singleURLs = types.Slice[string]{
		"https://music.163.com/#/song?id=1824625066",
		"https://music.163.com/song?id=2153793701&userid=757014",
	}

	unMatchURLs = types.Slice[string]{
		"https://www.zhihu.com/people/caiknife",
		"https://music.163.com/",
		"https://music.163.com/#/user/home?id=757014",
	}
)

func Test_regexp(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		assert.True(t, IsAlbumLink(s))
		v := albumRegexp.FindStringSubmatch(s)
		t.Log(len(v), v)
	})

	playlistURLs.ForEach(func(s string, i int) {
		assert.True(t, IsPlaylistLink(s))
		v := playlistRegexp.FindStringSubmatch(s)
		t.Log(len(v), v)
	})

	singleURLs.ForEach(func(s string, i int) {
		assert.True(t, IsSingleLink(s))
		v := singleRegexp.FindStringSubmatch(s)
		t.Log(len(v), v)
	})

	unMatchURLs.ForEach(func(s string, i int) {
		assert.False(t, IsSingleLink(s))
		assert.False(t, IsAlbumLink(s))
		assert.False(t, IsPlaylistLink(s))
	})
}

func TestSingleLinkID(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		id, err := SingleLinkID(s)
		assert.NoError(t, err)
		assert.True(t, id != 0)
	})
}

func TestAlbumLinkID(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		assert.NoError(t, err)
		assert.True(t, id != 0)
	})
}

func TestPlaylistLinkID(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		id, err := PlaylistLinkID(s)
		assert.NoError(t, err)
		assert.True(t, id != 0)
	})
}
