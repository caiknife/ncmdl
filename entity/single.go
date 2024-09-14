package entity

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

const nullSeparator = "\u0000"

type Single struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Artist types.Slice[*Artist] `json:"ar"`
	Album  *Album               `json:"al"`
	URL    string               `json:"url"`
}

func (s *Single) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}

func (s *Single) SavePath() string {
	return filepath.Join(s.Artist[0].Name, s.Album.Name)
}

func (s *Single) CoverURL() string {
	return s.Album.PicURL
}

func (s *Single) FileName() string {
	replaceName := strings.ReplaceAll(s.Name, "/", ",")
	return fmt.Sprintf("%s - %s.mp3", s.Artist[0].Name, replaceName)
}

func (s *Single) SaveFileName() string {
	return filepath.Join(s.SavePath(), s.FileName())
}

func (s *Single) AllArtistTag() string {
	allArtist := types.Slice[string]{}
	s.Artist.ForEach(func(artist *Artist, i int) {
		allArtist = append(allArtist, artist.Name)
	})
	return strings.Join(allArtist, nullSeparator)
}
