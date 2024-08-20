package main

import (
	"net/http"

	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
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
	Url        string      `json:"url"`
	ID         int         `json:"id"`
	CookieFile *CookieFile `json:"cookie_file"`
	DryRun     bool        `json:"dry_run"`
	Tmp        bool        `json:"tmp"`
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
	l = &Link{Url: url}

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

	err = l.id()
	if err != nil {
		err = errors.WithMessage(err, "link load ID")
		return nil, err
	}
	return l, nil
}

func (l *Link) id() (err error) {
	switch l.Type {
	case Single:
		l.ID, err = SingleLinkID(l.Url)
	case Album:
		l.ID, err = AlbumLinkID(l.Url)
	case Playlist:
		l.ID, err = PlaylistLinkID(l.Url)
	default:
		return ErrLinkIDNotMatch
	}
	if err != nil {
		err = errors.WithMessage(err, "parse link id")
		return err
	}
	return nil
}

func (l *Link) Download() (err error) {
	if l.DryRun {
		logger.ConsoleLogger.Infoln("当前是演习模式，要解析的URL是", l.Url)
		logger.ConsoleLogger.Infoln("退出程序，不进行下载")
		return nil
	}
	destDir := Path(lo.Ternary[string](l.Tmp, "./tmp/", "./"))
	switch l.Type {
	case Single:
		return DownloadSingle(l.ID, destDir)
	case Album:
		return DownloadAlbum(l.ID, destDir)
	case Playlist:
		return DownloadPlaylist(l.ID, destDir)
	}
	return nil
}

var (
	reqData = &utils.RequestData{}
)

func SetRequestDataCookie(cookie types.Slice[*http.Cookie]) {
	reqData.Cookies = cookie
}

func GetRequestData() *utils.RequestData {
	return reqData
}

type LinkOption func(*Link)

func OptionDryRun(dryRun bool) LinkOption {
	return func(l *Link) {
		l.DryRun = dryRun
	}
}

func OptionTmp(tmp bool) LinkOption {
	return func(l *Link) {
		l.Tmp = tmp
	}
}

func OptionCookieFile(cookieFile *CookieFile) LinkOption {
	return func(l *Link) {
		l.CookieFile = cookieFile
	}
}
