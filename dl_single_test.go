package main

import (
	"testing"
)

func TestSingleDetail(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		id, err := SingleLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := SingleDetail(id)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}

func TestDownloadSingle(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		id, err := SingleLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		err = DownloadSingle(id, "./tmp/")
		if err != nil {
			t.Error(err)
			return
		}
	})
}
