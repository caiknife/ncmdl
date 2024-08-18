package main

import (
	"net/http"

	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
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
}

func (l *Link) String() string {
	toString, _ := fjson.MarshalToString(l)
	return toString
}

const (
	ErrLinkTypeNotMatch types.Error = "链接不是歌单、专辑、单曲类型"
	ErrLinkIDNotMatch   types.Error = "链接无法匹配ID"
)

func NewLink(url string, cookieFile *CookieFile, dryRun bool) (l *Link, err error) {
	switch {
	case IsSingleLink(url):
		l = &Link{Type: Single, Url: url, CookieFile: cookieFile, DryRun: dryRun}
	case IsAlbumLink(url):
		l = &Link{Type: Album, Url: url, CookieFile: cookieFile, DryRun: dryRun}
	case IsPlaylistLink(url):
		l = &Link{Type: Playlist, Url: url, CookieFile: cookieFile, DryRun: dryRun}
	default:
		return nil, ErrLinkTypeNotMatch
	}
	err = l.id()
	if err != nil {
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
		return err
	}
	return nil
}

func (l *Link) Download() (err error) {
	if l.DryRun {
		logger.ConsoleLogger.Infoln("当前是演习模式，要解析的URL是", l.Url)
		logger.ConsoleLogger.Infoln("退出程序，不进行下载")
		return
	}
	switch l.Type {
	case Single:
		return DownloadSingle(l.ID, Path("./"))
	case Album:
		return DownloadAlbum(l.ID, Path("./"))
	case Playlist:
		return DownloadPlaylist(l.ID, Path("./"))
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
