package entity

import (
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
	return s.Artist[0].Name + "/" + s.Album.Name
}

func (s *Single) CoverURL() string {
	return s.Album.PicURL
}

func (s *Single) SaveFileName() string {
	return s.SavePath() + "/" + s.Artist[0].Name + " - " + s.Name + ".mp3"
}
