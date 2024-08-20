package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		v, err := NewLink(
			s,
			OptionCookieFile(NewCookieFile(ncmCookieFile)),
			OptionDryRun(false),
		)
		assert.NoError(t, err)
		t.Log(v)
	})
}
