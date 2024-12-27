package ncmdl

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
)

const (
	downloadResult = `{
    "data": [
        {
            "id": 1824625066,
            "url": "http://m801.music.126.net/20240817000252/37b18b2f62606413f469986606c4abd2/jdymusic/obj/wo3DlMOGwrbDjj7DisKw/7730680674/861c/4687/4748/bee7b941a70c1412b22daa08be9b0333.mp3",
            "br": 192000,
            "size": 5193604,
            "md5": "bee7b941a70c1412b22daa08be9b0333",
            "code": 200,
            "expi": 1200,
            "type": "mp3",
            "gain": -3.6819,
            "peak": 0,
            "fee": 1,
            "uf": null,
            "payed": 1,
            "flag": 196868,
            "canExtend": false,
            "freeTrialInfo": null,
            "level": "higher",
            "encodeType": "mp3",
            "channelLayout": null,
            "freeTrialPrivilege": {
                "resConsumable": false,
                "userConsumable": false,
                "listenType": null,
                "cannotListenReason": null,
                "playReason": null,
                "freeLimitTagType": null
            },
            "freeTimeTrialPrivilege": {
                "resConsumable": false,
                "userConsumable": false,
                "type": 0,
                "remainTime": 0
            },
            "urlSource": 0,
            "rightSource": 0,
            "podcastCtrp": null,
            "effectTypes": null,
            "time": 216321,
            "message": null,
            "levelConfuse": null
        }
    ],
    "code": 200
}`
)

func Test_DownloadResult(t *testing.T) {
	d := &DownloadResult{}
	err := fjson.UnmarshalFromString(downloadResult, d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(d)
}
