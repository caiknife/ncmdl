package ncmdl

import (
	"regexp"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/spf13/cast"
)

var (
	playlistRegexp = regexp.MustCompile(`^https://music\.163\.com/(#/)?playlist\?.*id=(\d+)`)
	albumRegexp    = regexp.MustCompile(`^https://music\.163\.com/(#/)?album\?.*id=(\d+)`)
	singleRegexp   = regexp.MustCompile(`^https://music\.163\.com/(#/)?song\?.*id=(\d+)`)
)

func IsSingleLink(s string) bool {
	return singleRegexp.MatchString(s)
}

func SingleLinkID(s string) (int, error) {
	submatch := types.Slice[string](singleRegexp.FindStringSubmatch(s))
	if submatch.IsEmpty() {
		return 0, ErrCannotFindSingleID
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
	return playlistRegexp.MatchString(s)
}

func PlaylistLinkID(s string) (int, error) {
	submatch := types.Slice[string](playlistRegexp.FindStringSubmatch(s))
	if submatch.IsEmpty() {
		return 0, ErrCannotFindPlaylistID
	}
	return cast.ToInt(submatch[2]), nil
}
