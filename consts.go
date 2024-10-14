package main

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	PoolSize      = 5
	NullSeparator = "\u0000"
)

const (
	ErrSongDetailIsEmpty types.Error = "单曲详情为空"
	ErrSongDownload      types.Error = "解析歌单歌曲下载链接异常"
)

const (
	ErrAlbumIsEmpty  types.Error = "专辑歌曲为空"
	ErrAlbumDownload types.Error = "解析专辑歌曲下载链接异常"
)

const (
	ErrPlaylistIsEmpty  types.Error = "歌单列表为空"
	ErrPlaylistDownload types.Error = "解析歌单歌曲下载链接异常"
)

const (
	ErrInputLinksAreEmpty types.Error = "请输入歌单、专辑、单曲链接"
)

const (
	ErrCookieFile types.Error = "cookies文件格式异常"
)

const (
	ErrCannotFindSingleID   types.Error = "没有匹配到单曲ID"
	ErrCannotFindAlbumID    types.Error = "没有匹配到专辑ID"
	ErrCannotFindPlaylistID types.Error = "没有匹配到歌单ID"
)
