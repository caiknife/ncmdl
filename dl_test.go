package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	t.Log(Path("."))
	t.Log(Path("./"))
	t.Log(Path("tmp"))
	t.Log(Path("tmp/"))
	t.Log(Path("./tmp"))
	t.Log(Path("../"))
}
