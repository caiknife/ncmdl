package ncmdl

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/stretchr/testify/assert"
)

var (
	albumURLs = types.Slice[string]{
		"https://music.163.com/album?id=195093556&userid=757014",
		"https://music.163.com/#/album?id=123837273",
		"https://music.163.com/#/album?app_version=9.2.25&id=39362733&dlt=0846",
	}

	playlistURLs = types.Slice[string]{
		"https://music.163.com/playlist?id=7350047928&userid=757014",
		"https://music.163.com/#/playlist?id=6623715587",
	}

	singleURLs = types.Slice[string]{
		"https://music.163.com/song?id=2153793701&userid=757014",
		"https://music.163.com/#/song?id=1824625066",
	}

	unMatchURLs = types.Slice[string]{
		"https://www.zhihu.com/people/caiknife",
		"https://music.163.com/",
		"https://music.163.com/#/user/home?id=757014",
	}
)

func TestIsAlbumLink(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		assert.Equal(t, true, IsAlbumLink(s))
	})
}

func TestIsPlaylistLink(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		assert.Equal(t, true, IsPlaylistLink(s))
	})
}

func TestIsSingleLink(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		assert.True(t, IsSingleLink(s))
	})
}

func TestNotMatch(t *testing.T) {
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
		t.Log(id)
	})
}

func TestAlbumLinkID(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		assert.NoError(t, err)
		assert.True(t, id != 0)
		t.Log(id)
	})
}

func TestPlaylistLinkID(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		id, err := PlaylistLinkID(s)
		assert.NoError(t, err)
		assert.True(t, id != 0)
		t.Log(id)
	})
}
