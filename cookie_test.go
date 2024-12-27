package ncmdl

import (
	"net/http"
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/stretchr/testify/assert"
)

const ncmCookieFile = "./ncm.txt"

func TestNewCookie(t *testing.T) {
	cookie := NewCookieFile(ncmCookieFile)
	AppLogger.Infoln("已经加载cookies文件", ncmCookieFile)
	assert.False(t, cookie.IsEmpty())
	cookie.Values.ForEach(func(cookie *http.Cookie, i int) {
		t.Log(i, cookie)
	})
}

type testStruct struct {
	types.Slice[string]
}

func (t *testStruct) String() string {
	toString, _ := fjson.MarshalToString(t)
	return toString
}

func TestSlice(t *testing.T) {
	sss := &testStruct{
		Slice: []string{
			"1", "2", "3", "4", "5",
		},
	}
	t.Log(sss)

	sss.ForEach(func(s string, i int) {
		t.Log(i, s)
	})
}
