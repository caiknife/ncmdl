package main

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
)

const (
	singleResult = `{
    "songs": [
        {
            "name": "Squatty Roo",
            "id": 1824625066,
            "pst": 0,
            "t": 0,
            "ar": [
                {
                    "id": 47024056,
                    "name": "Enric Peidro & Jonathan Stout",
                    "tns": [],
                    "alias": []
                }
            ],
            "alia": [],
            "pop": 15,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 6,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p2.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 216321,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 8655978,
                "vd": 568,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 5193604,
                "vd": 3180,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 3462417,
                "vd": 4964,
                "sr": 44100
            },
            "sq": {
                "br": 1601567,
                "fid": 0,
                "size": 43306631,
                "vd": 574,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 4,
            "rtUrl": null,
            "ftype": 0,
            "rtUrls": [],
            "djId": 0,
            "copyright": 1,
            "s_id": 0,
            "mark": 270336,
            "originCoverType": 0,
            "originSongSimpleData": null,
            "tagPicList": null,
            "resourceState": true,
            "version": 6,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "mv": 0,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "publishTime": 1615478400000
        }
    ],
    "privileges": [
        {
            "id": 1824625066,
            "fee": 1,
            "payed": 1,
            "st": 0,
            "pl": 999000,
            "dl": 999000,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 0,
            "toast": false,
            "flag": 196868,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "jymaster",
            "playMaxBrLevel": "jymaster",
            "downloadMaxBrLevel": "jymaster",
            "plLevel": "jymaster",
            "dlLevel": "jymaster",
            "flLevel": "none",
            "rscl": null,
            "freeTrialPrivilege": {
                "resConsumable": false,
                "userConsumable": false,
                "listenType": null,
                "cannotListenReason": null,
                "playReason": null,
                "freeLimitTagType": null
            },
            "rightSource": 0,
            "chargeInfoList": [
                {
                    "rate": 128000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 1
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 1
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 1
                },
                {
                    "rate": 999000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 1
                }
            ],
            "code": 0,
            "message": null
        }
    ],
    "code": 200
}`
)

func Test_SingleResult(t *testing.T) {
	d := &SingleResult{}
	err := fjson.UnmarshalFromString(singleResult, d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(d)
}

func Test_SingleDetail(t *testing.T) {
	singleURLs.ForEach(func(s string, i int) {
		id, err := SingleLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := SingleDetail(id, reqData)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}
