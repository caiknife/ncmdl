package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLink_Single(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		v, err := NewLink(
			s,
			LinkOptionCookieFile(NewCookieFile(ncmCookieFile)),
			LinkOptionOnlyInfo(true),
			LinkOptionTmp(true),
		)
		assert.NoError(t, err)
		err = v.Download()
		if err != nil {
			t.Error(err)
			return
		}
	})
}

func Test_NewLink_Album(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		v, err := NewLink(
			s,
			LinkOptionCookieFile(NewCookieFile(ncmCookieFile)),
			LinkOptionOnlyInfo(true),
			LinkOptionTmp(true),
		)
		assert.NoError(t, err)
		err = v.Download()
		if err != nil {
			t.Error(err)
			return
		}
	})
}
