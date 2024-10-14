package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/cast"
)

type CookieFile struct {
	Values types.Slice[*http.Cookie] `json:"values"`
}

func (c *CookieFile) String() string {
	toString, _ := fjson.MarshalToString(c)
	return toString
}

func (c *CookieFile) IsEmpty() bool {
	return c.Values.IsEmpty()
}

func (c *CookieFile) ToHttpCookie() (cookie types.Slice[*http.Cookie]) {
	if c.IsEmpty() {
		return cookie
	}
	return c.Values
}

const (
	ErrCookieFile = "cookies文件格式异常"
)

func NewCookieFile(cookieFile string) *CookieFile {
	c := &CookieFile{
		Values: types.Slice[*http.Cookie]{},
	}
	if cookieFile == "" {
		return c
	}
	line, err := fileutil.ReadFileByLine(cookieFile)
	if err != nil {
		return c
	}
	for _, s := range line {
		if strings.HasPrefix(s, "#") || s == "" {
			continue
		}
		split := strings.Split(s, "\t")
		if len(split) != 7 {
			panic(ErrCookieFile)
		}

		c.Values = append(c.Values, &http.Cookie{
			Name:     split[5],
			Value:    split[6],
			Path:     split[2],
			Domain:   split[0],
			Expires:  time.Unix(int64(cast.ToInt(split[4])), 0),
			MaxAge:   cast.ToInt(split[4]),
			Secure:   split[3] == "TRUE",
			SameSite: http.SameSiteLaxMode,
		})
	}
	return c
}
