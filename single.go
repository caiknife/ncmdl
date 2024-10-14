package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
)

type SingleInfo struct {
	ID     int                      `json:"id"`
	Name   string                   `json:"name"`
	Artist types.Slice[*ArtistInfo] `json:"ar"`
	Album  *AlbumInfo               `json:"al"`
	URL    string                   `json:"url"`
}

func (s *SingleInfo) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}

func (s *SingleInfo) SavePath() string {
	return filepath.Join(s.Artist[0].Name, s.Album.Name)
}

func (s *SingleInfo) CoverURL() string {
	return s.Album.PicURL
}

func (s *SingleInfo) FileName() string {
	replaceName := strings.ReplaceAll(s.Name, "/", ",")
	return fmt.Sprintf("%s - %s.mp3", s.Artist[0].Name, replaceName)
}

func (s *SingleInfo) SaveFileName() string {
	return filepath.Join(s.SavePath(), s.FileName())
}

func (s *SingleInfo) AllArtistsTag() string {
	allArtists := types.Slice[string]{}
	s.Artist.ForEach(func(info *ArtistInfo, i int) {
		allArtists = append(allArtists, info.Name)
	})
	return strings.Join(allArtists, NullSeparator)
}

type SingleResult struct {
	Songs types.Slice[*SingleInfo] `json:"songs"`
}

func (s *SingleResult) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}

func SingleDetail(singleID int, reqData *utils.RequestData) (result types.Slice[*SingleInfo], err error) {
	detail, err := api.GetSongDetail(*reqData, []int{singleID})
	if err != nil {
		err = errors.WithMessage(err, "api get song detail")
		return nil, err
	}
	d := &SingleResult{}
	err = fjson.UnmarshalFromString(detail.RawJson, d)
	if err != nil {
		err = errors.WithMessage(err, "json unmarshal")
		return nil, err
	}
	if d.Songs.IsEmpty() {
		return nil, ErrSongDetailIsEmpty
	}
	return d.Songs, nil
}
