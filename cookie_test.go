package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ncmCookieFile = "./ncm.txt"

func TestNewCookie(t *testing.T) {
	cookie := NewCookieFile(ncmCookieFile)
	assert.False(t, cookie.IsEmpty())
	t.Log(cookie.ToHttpCookie())
}
