package main

import (
	"net/http"
	"testing"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/caiknife/mp3lister/lib/types"
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

func Test_song_url(t *testing.T) {
	url, err := api.GetSongURL(*GetRequestData(), api.SongURLConfig{
		EncodeType: "",
		Level:      "higher",
		Ids: []int{
			1824625066,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url.RawJson)
}

func Test_song_detail(t *testing.T) {
	reqData = &utils.RequestData{
		Cookies: []*http.Cookie{
			{
				Name:  "MUSIC_U", // 获取无损音质需填写 Cookies 中的 MUSIC_U
				Value: "00561BD9B9C176029618578602A07EDAB00B6C7187801A3AC5D754B86F4607615A24A1F4BECAB0867FFFE453A9B68CAD22C5CC5E25CF9676F65B1AC0032E23939C57307100D03CEE3BDF4054FB40E2DC4D9BA5048AAFED08B08F8F7C420C2BC086E8E0DE51F789713D387782765011502EBE3561C26BE6B20F85B5C532115D191D64126DEED75999C5DFF4DD4E00AACC68DE928B4C6F2B55C571165FF16393A095F9A3EB5F47B336CC0448CD64E4591BF0F59208714EA0132CC57B6739FB0F923CD1E38C693E0033DD4904C7ECB84AE5638677957B1A3EA1B1D01206DC652E081B343E2D723BBD9553019973CFBC3A481BB0C405CDBAC0B42A21F3071CC9F0F51BEC2FA1569ACDECEE032532C6E027F5AB4FC04448CA447A45F291920219AADB99974B92F65CFE356A99CA95EDCFEDD603CAC7BE38FAF1CA731FDEE4E3C4D55D79",
			},
		},
	}
	detail, err := api.GetSongDetail(*reqData, types.Slice[int]{1824625066})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(detail.RawJson)
}
