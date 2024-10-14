package main

import (
	"path/filepath"

	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type LinkType int

const (
	Single LinkType = iota + 1
	Album
	Playlist
)

type Link struct {
	Type       LinkType    `json:"type"`
	URL        string      `json:"url"`
	ID         int         `json:"id"`
	CookieFile *CookieFile `json:"cookie_file"`
	Info       bool        `json:"info"`
	Tmp        bool        `json:"tmp"`

	reqData *utils.RequestData
	destDir string
}

func (l *Link) String() string {
	toString, _ := fjson.MarshalToString(l)
	return toString
}

const (
	ErrLinkTypeNotMatch types.Error = "链接不是歌单、专辑、单曲类型"
	ErrLinkIDNotMatch   types.Error = "链接无法匹配ID"
)

func NewLink(url string, opts ...LinkOption) (l *Link, err error) {
	l = &Link{URL: url}

	switch {
	case IsSingleLink(url):
		l.Type = Single
	case IsAlbumLink(url):
		l.Type = Album
	case IsPlaylistLink(url):
		l.Type = Playlist
	default:
		return nil, ErrLinkTypeNotMatch
	}

	for _, opt := range opts {
		opt(l)
	}

	err = l.loadID()
	if err != nil {
		err = errors.WithMessage(err, "link load ID")
		return nil, err
	}

	l.reqData = &utils.RequestData{}
	l.reqData.Cookies = l.CookieFile.ToHttpCookie()

	l.destDir = Path(lo.Ternary[string](l.Tmp, filepath.Join(".", "tmp"), "."))

	return l, nil
}

func (l *Link) Download() error {
	switch l.Type {
	case Single:
		return l.downloadSingle()
	case Album:
		return l.downloadAlbum()
	case Playlist:
		return l.downloadPlaylist()
	default:
		return ErrLinkTypeNotMatch
	}
	return nil
}

func (l *Link) downloadSingle() error {
	return nil
}

func (l *Link) downloadAlbum() error {
	return nil
}

func (l *Link) downloadPlaylist() error {
	return nil
}

func (l *Link) loadID() (err error) {
	switch l.Type {
	case Single:
		l.ID, err = SingleLinkID(l.URL)
	case Album:
		l.ID, err = AlbumLinkID(l.URL)
	case Playlist:
		l.ID, err = PlaylistLinkID(l.URL)
	default:
		return ErrLinkIDNotMatch
	}
	if err != nil {
		err = errors.WithMessage(err, "parse link id")
		return err
	}
	return nil
}

type LinkOption func(link *Link)

func LinkOptionOnlyInfo(flag bool) LinkOption {
	return func(link *Link) {
		link.Info = flag
	}
}

func LinkOptionTmp(flag bool) LinkOption {
	return func(link *Link) {
		link.Tmp = flag
	}
}

func LinkOptionCookieFile(file *CookieFile) LinkOption {
	return func(link *Link) {
		link.CookieFile = file
	}
}
