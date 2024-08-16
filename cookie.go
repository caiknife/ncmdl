package main

import (
	"net/http"
	"strings"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/samber/lo"
)

type CookieFile struct {
	Values types.Map[string] `json:"values"`
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
	cookie = lo.MapToSlice[string, string, *http.Cookie](c.Values, func(key string, value string) *http.Cookie {
		return &http.Cookie{
			Name:  key,
			Value: value,
		}
	})
	return cookie
}

func NewCookieFile(cookieFile string) *CookieFile {
	c := &CookieFile{
		Values: types.Map[string]{},
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
		c.Values[split[5]] = split[6]
	}
	return c
}
