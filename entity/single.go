package entity

import (
	"fmt"
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Single struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Artist types.Slice[*Artist] `json:"ar"`
	Album  *Album               `json:"al"`
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
	return fmt.Sprintf("%s - %s.mp3", s.Artist[0].Name, s.Name)
}

func (s *Single) SaveFileName() string {
	return filepath.Join(s.SavePath(), s.FileName())
}
