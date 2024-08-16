package main

import (
	"regexp"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/spf13/cast"
)

var (
	albumRegexp    = regexp.MustCompile(`^https://music.163.com/(#/)?album\?id=(\d+)`)
	playlistRegexp = regexp.MustCompile(`^https://music.163.com/(#/)?playlist\?id=(\d+)`)
	singleRegexp   = regexp.MustCompile(`^https://music.163.com/(#/)?song\?id=(\d+)`)
)

const (
	ErrCannotFindSingleLinkID types.Error = "没有匹配到单曲ID"
	ErrCannotFindAlbumID      types.Error = "没有匹配到专辑ID"
	ErrCannotFindPlaylist     types.Error = "没有匹配到歌单ID"
)

func IsSingleLink(s string) bool {
	return singleRegexp.MatchString(s)
}

func SingleLinkID(s string) (int, error) {
	submatch := types.Slice[string](singleRegexp.FindStringSubmatch(s))
	if submatch.IsEmpty() {
		return 0, ErrCannotFindSingleLinkID
	}
	return cast.ToInt(submatch[2]), nil
}

func IsAlbumLink(s string) bool {
	return albumRegexp.MatchString(s)
}

func AlbumLinkID(s string) (int, error) {
	submatch := types.Slice[string](albumRegexp.FindStringSubmatch(s))
	if submatch.IsEmpty() {
		return 0, ErrCannotFindAlbumID
	}
	return cast.ToInt(submatch[2]), nil
}

func IsPlaylistLink(s string) bool {
	return albumRegexp.MatchString(s)
}

func PlaylistLinkID(s string) (int, error) {
	submatch := types.Slice[string](playlistRegexp.FindStringSubmatch(s))
	if submatch.IsEmpty() {
		return 0, ErrCannotFindPlaylist
	}
	return cast.ToInt(submatch[2]), nil
}
