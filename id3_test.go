package main

import (
	"strings"
	"testing"

	"github.com/bogem/id3v2/v2"
)

const mp3File = "./tmp/Enric Peidro & Jonathan Stout/Groove at First Sight/Enric Peidro & Jonathan Stout - Squatty Roo.mp3"
const mp3File_1 = "./tmp/Jonathan Stout and his Campus Five/Jammin' the Blues/Jonathan Stout and his Campus Five - Diga Diga Doo.mp3"

func TestID3_GetArtist(t *testing.T) {
	open, err := id3v2.Open(mp3File_1, id3v2.Options{Parse: true})
	if err != nil {
		t.Error(err)
		return
	}
	defer open.Close()

	t.Log(open.Artist())
}

func TestID3_SetArtist(t *testing.T) {
	open, err := id3v2.Open(mp3File, id3v2.Options{Parse: true})
	if err != nil {
		t.Error(err)
		return
	}
	defer open.Close()

	open.SetArtist(strings.Join([]string{"111", "222", "333"}, NullSeparator))
	err = open.Save()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(open.Artist())
}
