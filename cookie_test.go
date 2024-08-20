package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ncmCookieFile = "./ncm.txt"

func init() {
	cookie := NewCookieFile(ncmCookieFile)
	SetRequestDataCookie(cookie.ToHttpCookie())
}
func TestNewCookie(t *testing.T) {
	cookie := NewCookieFile(ncmCookieFile)
	assert.False(t, cookie.IsEmpty())
	cookie.Values.ForEach(func(cookie *http.Cookie, i int) {
		t.Log(i, cookie)
	})
}
