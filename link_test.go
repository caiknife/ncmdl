package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		v, err := NewLink(s, NewCookieFile(ncmCookieFile))
		assert.NoError(t, err)
		t.Log(v)
	})
}
