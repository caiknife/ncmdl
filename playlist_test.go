package main

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
)

const (
	playlistResult = `{
    "code": 200,
    "relatedVideos": {},
    "playlist": {
        "id": 6623715587,
        "name": "Swing with Riffs",
        "coverImgId": 109951164916554400,
        "coverImgUrl": "https://p1.music.126.net/RnaxjHJPsXGZyOZBNYfHnw==/109951164916554394.jpg",
        "coverImgId_str": "109951164916554394",
        "adType": 0,
        "userId": 757014,
        "createTime": 1613714408316,
        "status": 0,
        "opRecommend": false,
        "highQuality": false,
        "newImported": false,
        "updateTime": 1614999988857,
        "trackCount": 12,
        "specialType": 0,
        "privacy": 0,
        "trackUpdateTime": 1719296032660,
        "commentThreadId": "A_PL_0_6623715587",
        "playCount": 152,
        "trackNumberUpdateTime": 1614999988857,
        "subscribedCount": 13,
        "cloudTrackCount": 0,
        "ordered": false,
        "description": null,
        "tags": [],
        "updateFrequency": null,
        "backgroundCoverId": 0,
        "backgroundCoverUrl": null,
        "titleImage": 0,
        "titleImageUrl": null,
        "detailPageTitle": null,
        "englishTitle": null,
        "officialPlaylistType": null,
        "copied": false,
        "relateResType": null,
        "coverStatus": 1,
        "subscribers": [
            {
                "defaultAvatar": false,
                "province": 320000,
                "authStatus": 0,
                "followed": false,
                "avatarUrl": "http://p1.music.126.net/wYvZdEMSM3uwWYwajx-2Ig==/109951169836144188.jpg",
                "accountStatus": 0,
                "gender": 1,
                "city": 320100,
                "birthday": 0,
                "userId": 256569466,
                "userType": 0,
                "nickname": "去游泳",
                "signature": "去游泳",
                "description": "",
                "detailDescription": "",
                "avatarImgId": 109951169836144200,
                "backgroundImgId": 2002210674180200,
                "backgroundUrl": "http://p1.music.126.net/45Nu4EqvFqK_kQj6BkPwcw==/2002210674180200.jpg",
                "authority": 0,
                "mutual": false,
                "expertTags": null,
                "experts": null,
                "djStatus": 0,
                "vipType": 11,
                "remarkName": null,
                "authenticationTypes": 0,
                "avatarDetail": null,
                "backgroundImgIdStr": "2002210674180200",
                "avatarImgIdStr": "109951169836144188",
                "anchor": false,
                "avatarImgId_str": "109951169836144188"
            },
            {
                "defaultAvatar": false,
                "province": 430000,
                "authStatus": 0,
                "followed": false,
                "avatarUrl": "http://p1.music.126.net/E_pLhjC9JL_LKBwJnugTZQ==/109951164805988735.jpg",
                "accountStatus": 0,
                "gender": 1,
                "city": 430100,
                "birthday": 0,
                "userId": 447745015,
                "userType": 0,
                "nickname": "天叔uncletian",
                "signature": "",
                "description": "",
                "detailDescription": "",
                "avatarImgId": 109951164805988740,
                "backgroundImgId": 109951162868128400,
                "backgroundUrl": "http://p1.music.126.net/2zSNIqTcpHL2jIvU6hG0EA==/109951162868128395.jpg",
                "authority": 0,
                "mutual": false,
                "expertTags": null,
                "experts": null,
                "djStatus": 0,
                "vipType": 11,
                "remarkName": null,
                "authenticationTypes": 0,
                "avatarDetail": null,
                "backgroundImgIdStr": "109951162868128395",
                "avatarImgIdStr": "109951164805988735",
                "anchor": false,
                "avatarImgId_str": "109951164805988735"
            },
            {
                "defaultAvatar": false,
                "province": 430000,
                "authStatus": 0,
                "followed": false,
                "avatarUrl": "http://p1.music.126.net/LAd7gGASYGwsH7jK1oGnrA==/109951165624406169.jpg",
                "accountStatus": 0,
                "gender": 1,
                "city": 430100,
                "birthday": 0,
                "userId": 1515402008,
                "userType": 0,
                "nickname": "岳麓LouisZhu",
                "signature": "",
                "description": "",
                "detailDescription": "",
                "avatarImgId": 109951165624406180,
                "backgroundImgId": 109951162868126480,
                "backgroundUrl": "http://p1.music.126.net/_f8R60U9mZ42sSNvdPn2sQ==/109951162868126486.jpg",
                "authority": 0,
                "mutual": false,
                "expertTags": null,
                "experts": null,
                "djStatus": 0,
                "vipType": 11,
                "remarkName": null,
                "authenticationTypes": 0,
                "avatarDetail": null,
                "backgroundImgIdStr": "109951162868126486",
                "avatarImgIdStr": "109951165624406169",
                "anchor": false,
                "avatarImgId_str": "109951165624406169"
            },
            {
                "defaultAvatar": false,
                "province": 440000,
                "authStatus": 0,
                "followed": false,
                "avatarUrl": "http://p1.music.126.net/cyMdrjRMafrBSa5I5lW8hw==/109951162845906381.jpg",
                "accountStatus": 0,
                "gender": 1,
                "city": 440300,
                "birthday": 0,
                "userId": 105183191,
                "userType": 0,
                "nickname": "金鱼花火o",
                "signature": "",
                "description": "",
                "detailDescription": "",
                "avatarImgId": 109951162845906380,
                "backgroundImgId": 2002210674180203,
                "backgroundUrl": "http://p1.music.126.net/bmA_ablsXpq3Tk9HlEg9sA==/2002210674180203.jpg",
                "authority": 0,
                "mutual": false,
                "expertTags": null,
                "experts": null,
                "djStatus": 0,
                "vipType": 11,
                "remarkName": null,
                "authenticationTypes": 0,
                "avatarDetail": null,
                "backgroundImgIdStr": "2002210674180203",
                "avatarImgIdStr": "109951162845906381",
                "anchor": false,
                "avatarImgId_str": "109951162845906381"
            },
            {
                "defaultAvatar": false,
                "province": 430000,
                "authStatus": 0,
                "followed": false,
                "avatarUrl": "http://p1.music.126.net/ref34rK2fmSmqmzMSx7b7w==/109951164343197348.jpg",
                "accountStatus": 0,
                "gender": 1,
                "city": 430100,
                "birthday": 0,
                "userId": 330160148,
                "userType": 0,
                "nickname": "ju胖胖养殖场ceo",
                "signature": "",
                "description": "",
                "detailDescription": "",
                "avatarImgId": 109951164343197340,
                "backgroundImgId": 2002210674180200,
                "backgroundUrl": "http://p1.music.126.net/45Nu4EqvFqK_kQj6BkPwcw==/2002210674180200.jpg",
                "authority": 0,
                "mutual": false,
                "expertTags": null,
                "experts": null,
                "djStatus": 0,
                "vipType": 11,
                "remarkName": null,
                "authenticationTypes": 0,
                "avatarDetail": null,
                "backgroundImgIdStr": "2002210674180200",
                "avatarImgIdStr": "109951164343197348",
                "anchor": false,
                "avatarImgId_str": "109951164343197348"
            }
        ],
        "subscribed": null,
        "creator": {
            "defaultAvatar": false,
            "province": 430000,
            "authStatus": 0,
            "followed": false,
            "avatarUrl": "http://p1.music.126.net/W0clmVaZY6W4YN23TMZlgw==/18546562139165524.jpg",
            "accountStatus": 0,
            "gender": 1,
            "city": 430100,
            "birthday": 0,
            "userId": 757014,
            "userType": 0,
            "nickname": "efinKiaC",
            "signature": "今天你流的泪，都是当初没流的汗。",
            "description": "",
            "detailDescription": "",
            "avatarImgId": 18546562139165524,
            "backgroundImgId": 109951163560892770,
            "backgroundUrl": "http://p1.music.126.net/pDsDPWmb5Oa3p_XfxwBUjw==/109951163560892763.jpg",
            "authority": 0,
            "mutual": false,
            "expertTags": null,
            "experts": null,
            "djStatus": 10,
            "vipType": 11,
            "remarkName": null,
            "authenticationTypes": 0,
            "avatarDetail": null,
            "backgroundImgIdStr": "109951163560892763",
            "avatarImgIdStr": "18546562139165524",
            "anchor": false,
            "avatarImgId_str": "18546562139165524"
        },
        "tracks": [
            {
                "name": "C Jam Blues",
                "id": 1441524915,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 0,
                        "name": "Duke Ellington Big Band",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 5,
                "st": 0,
                "rt": "",
                "fee": 8,
                "v": 6,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 88194564,
                    "name": "C Jam Blues",
                    "picUrl": "http://p1.music.126.net/RnaxjHJPsXGZyOZBNYfHnw==/109951164916554394.jpg",
                    "tns": [],
                    "pic_str": "109951164916554394",
                    "pic": 109951164916554400
                },
                "dt": 294120,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 11767685,
                    "vd": 18940,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 7060628,
                    "vd": 21512,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 4707100,
                    "vd": 23095,
                    "sr": 44100
                },
                "sq": {
                    "br": 1067000,
                    "fid": 0,
                    "size": 31881455,
                    "vd": 18947,
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
                "mark": 139392,
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
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 743010,
                "mv": 0,
                "publishTime": 1585238400000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Tuxedo Junction",
                "id": 1326804162,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 49004,
                        "name": "BBC Big Band",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 5,
                "st": 0,
                "rt": "",
                "fee": 8,
                "v": 4,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 74480243,
                    "name": "经典大乐团",
                    "picUrl": "http://p1.music.126.net/GZTVu5PGXm4Cfy7j96nHOQ==/109951163676539342.jpg",
                    "tns": [],
                    "pic_str": "109951163676539342",
                    "pic": 109951163676539340
                },
                "dt": 206245,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 8252648,
                    "vd": 0,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 4951606,
                    "vd": 0,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 3301085,
                    "vd": 0,
                    "sr": 44100
                },
                "sq": {
                    "br": 733388,
                    "fid": 0,
                    "size": 18907285,
                    "vd": 0,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 15,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 270464,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 4,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 743010,
                "mv": 0,
                "publishTime": 1495555200000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Take the A Train",
                "id": 1326804169,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 49004,
                        "name": "BBC Big Band",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 8,
                "v": 4,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 74480243,
                    "name": "经典大乐团",
                    "picUrl": "http://p1.music.126.net/GZTVu5PGXm4Cfy7j96nHOQ==/109951163676539342.jpg",
                    "tns": [],
                    "pic_str": "109951163676539342",
                    "pic": 109951163676539340
                },
                "dt": 202274,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 8093823,
                    "vd": 1,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 4856311,
                    "vd": 1,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 3237555,
                    "vd": 1,
                    "sr": 44100
                },
                "sq": {
                    "br": 811480,
                    "fid": 0,
                    "size": 20517761,
                    "vd": 1,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 24,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 270464,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 4,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 743010,
                "mv": 0,
                "publishTime": 1495555200000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "The Daly Jump",
                "id": 1375171522,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 29999,
                        "name": "Count Basie",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 0,
                "v": 9,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 80110184,
                    "name": "Not Now, I'll Tell You When (Remastered Version)",
                    "picUrl": "http://p1.music.126.net/dYf0C80S1FitAL4zgIksNQ==/109951164185404676.jpg",
                    "tns": [],
                    "pic_str": "109951164185404676",
                    "pic": 109951164185404670
                },
                "dt": 250881,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 10038378,
                    "vd": -30766,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 6023044,
                    "vd": -28240,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 4015377,
                    "vd": -26948,
                    "sr": 44100
                },
                "sq": {
                    "br": 1411000,
                    "fid": 0,
                    "size": 29874851,
                    "vd": -30749,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 6,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 262144,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 9,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 743010,
                "mv": 0,
                "publishTime": 1527264000000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Them There Eyes",
                "id": 1405103306,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 508416,
                        "name": "James Morrison",
                        "tns": [],
                        "alias": []
                    },
                    {
                        "id": 54713,
                        "name": "Deni Hines",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 8,
                "v": 4,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 83567739,
                    "name": "The Other Woman",
                    "picUrl": "http://p1.music.126.net/woj-IPvYiB4psGnstFV9Kw==/109951164501469328.jpg",
                    "tns": [],
                    "pic_str": "109951164501469328",
                    "pic": 109951164501469330
                },
                "dt": 186106,
                "h": {
                    "br": 320002,
                    "fid": 0,
                    "size": 7447031,
                    "vd": -13119,
                    "sr": 44100
                },
                "m": {
                    "br": 192002,
                    "fid": 0,
                    "size": 4468236,
                    "vd": -10543,
                    "sr": 44100
                },
                "l": {
                    "br": 128002,
                    "fid": 0,
                    "size": 2978839,
                    "vd": -8829,
                    "sr": 44100
                },
                "sq": {
                    "br": 914204,
                    "fid": 0,
                    "size": 21267446,
                    "vd": -14072,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 6,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 8192,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 4,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 7003,
                "mv": 0,
                "publishTime": 1167580800000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Little Brown Jug",
                "id": 535709086,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 317399,
                        "name": "SWR Big Band",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": null,
                "fee": 8,
                "v": 7,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 37516944,
                    "name": "Glenn Miller - Die größten Erfolge",
                    "picUrl": "http://p1.music.126.net/JqPKZ2Y5Ugn5Jz1zm7DAzg==/109951163132711115.jpg",
                    "tns": [],
                    "pic_str": "109951163132711115",
                    "pic": 109951163132711120
                },
                "dt": 169160,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 6768893,
                    "vd": 8022,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 4061353,
                    "vd": 11272,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 2707583,
                    "vd": 9258,
                    "sr": 44100
                },
                "sq": {
                    "br": 848191,
                    "fid": 0,
                    "size": 17934999,
                    "vd": 8217,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 6,
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
                "version": 7,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 456010,
                "mv": 0,
                "publishTime": 1372089600007,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Tuxedo Junction",
                "id": 1338791509,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 31222,
                        "name": "Duke Ellington",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 8,
                "v": 11,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 75122371,
                    "name": "Will Big Bands Ever Come Back (HD Remastered)",
                    "picUrl": "http://p1.music.126.net/oHp6hs9W1bLd_arYZNyFzg==/109951163786381986.jpg",
                    "tns": [],
                    "pic_str": "109951163786381986",
                    "pic": 109951163786381980
                },
                "dt": 209214,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 8370721,
                    "vd": 0,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 5022450,
                    "vd": 0,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 3348315,
                    "vd": 0,
                    "sr": 44100
                },
                "sq": {
                    "br": 927014,
                    "fid": 0,
                    "size": 24243119,
                    "vd": 0,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 1,
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
                "version": 11,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 456010,
                "mv": 0,
                "publishTime": 1541520000007,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "In the Mood",
                "id": 535709081,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 317399,
                        "name": "SWR Big Band",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": null,
                "fee": 8,
                "v": 9,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 37516944,
                    "name": "Glenn Miller - Die größten Erfolge",
                    "picUrl": "http://p1.music.126.net/JqPKZ2Y5Ugn5Jz1zm7DAzg==/109951163132711115.jpg",
                    "tns": [],
                    "pic_str": "109951163132711115",
                    "pic": 109951163132711120
                },
                "dt": 215200,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 8611048,
                    "vd": 1246,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 5166646,
                    "vd": 2334,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 3444445,
                    "vd": 4680,
                    "sr": 44100
                },
                "sq": {
                    "br": 847016,
                    "fid": 0,
                    "size": 22784746,
                    "vd": 1761,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 1,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 139264,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 9,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 456010,
                "mv": 0,
                "publishTime": 1372089600007,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Lullaby of Birdland",
                "id": 22788309,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 16664,
                        "name": "JUJU",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 1,
                "v": 35,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 2094121,
                    "name": "DELICIOUS",
                    "picUrl": "http://p1.music.126.net/xIKx1xGc_rXOigBcRDUaIw==/109951166198270435.jpg",
                    "tns": [],
                    "pic_str": "109951166198270435",
                    "pic": 109951166198270430
                },
                "dt": 155759,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 6232860,
                    "vd": -27673,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 3739733,
                    "vd": -25052,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 2493170,
                    "vd": -23313,
                    "sr": 44100
                },
                "sq": {
                    "br": 1533798,
                    "fid": 0,
                    "size": 29862943,
                    "vd": -28322,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "1",
                "no": 7,
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
                "version": 35,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 2706476,
                "mv": 28858,
                "publishTime": 1322582400000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": {
                        "vid": "28858",
                        "type": 0,
                        "title": "Lullaby Of Birdland",
                        "playTime": 11327,
                        "coverUrl": "http://p2.music.126.net/71z7PbFp5Mo42HueZhBVHA==/109951166536568401.jpg",
                        "publishTime": 1368633600007,
                        "artists": null,
                        "alias": null
                    }
                }
            },
            {
                "name": "Satin Doll",
                "id": 464714353,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 12545080,
                        "name": "Joe Eckert",
                        "tns": [],
                        "alias": []
                    },
                    {
                        "id": 31101420,
                        "name": "United States Air Force Band - Airmen of Note",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": null,
                "fee": 8,
                "v": 13,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 35265367,
                    "name": "UNITED STATES AIR FORCE AIRMEN OF NOTE: Let's Dance",
                    "picUrl": "http://p1.music.126.net/a7qEOUlgBDVL_6IAkLHD4g==/18154036487011510.jpg",
                    "tns": [],
                    "pic_str": "18154036487011510",
                    "pic": 18154036487011510
                },
                "dt": 309960,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 12400893,
                    "vd": -8000,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 7440553,
                    "vd": -5400,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 4960383,
                    "vd": -4000,
                    "sr": 44100
                },
                "sq": {
                    "br": 806966,
                    "fid": 0,
                    "size": 31265904,
                    "vd": -8000,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 18,
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
                "version": 13,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 471010,
                "mv": 0,
                "publishTime": 1310486400007,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Opus No. 1",
                "id": 464714348,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 31101420,
                        "name": "United States Air Force Band - Airmen of Note",
                        "tns": [],
                        "alias": []
                    },
                    {
                        "id": 12545080,
                        "name": "Joe Eckert",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": null,
                "fee": 8,
                "v": 13,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 35265367,
                    "name": "UNITED STATES AIR FORCE AIRMEN OF NOTE: Let's Dance",
                    "picUrl": "http://p1.music.126.net/a7qEOUlgBDVL_6IAkLHD4g==/18154036487011510.jpg",
                    "tns": [],
                    "pic_str": "18154036487011510",
                    "pic": 18154036487011510
                },
                "dt": 180120,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 7207750,
                    "vd": -19600,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 4324667,
                    "vd": -17100,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 2883126,
                    "vd": -15500,
                    "sr": 44100
                },
                "sq": {
                    "br": 949724,
                    "fid": 0,
                    "size": 21383056,
                    "vd": -19300,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "01",
                "no": 13,
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
                "version": 13,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 471010,
                "mv": 0,
                "publishTime": 1310486400007,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            },
            {
                "name": "Flat Foot Floogie",
                "id": 22680695,
                "pst": 0,
                "t": 0,
                "ar": [
                    {
                        "id": 15993,
                        "name": "akiko",
                        "tns": [],
                        "alias": []
                    }
                ],
                "alia": [],
                "pop": 10,
                "st": 0,
                "rt": "",
                "fee": 1,
                "v": 15,
                "crbt": null,
                "cf": "",
                "al": {
                    "id": 2080626,
                    "name": "Little Miss Jazz and Jive Goes Around the World",
                    "picUrl": "http://p1.music.126.net/4fyo1hQzms8S3y_K89GwHg==/109951167465449309.jpg",
                    "tns": [],
                    "pic_str": "109951167465449309",
                    "pic": 109951167465449310
                },
                "dt": 207818,
                "h": {
                    "br": 320000,
                    "fid": 0,
                    "size": 8315342,
                    "vd": -45642,
                    "sr": 44100
                },
                "m": {
                    "br": 192000,
                    "fid": 0,
                    "size": 4989222,
                    "vd": -43003,
                    "sr": 44100
                },
                "l": {
                    "br": 128000,
                    "fid": 0,
                    "size": 3326163,
                    "vd": -41207,
                    "sr": 44100
                },
                "sq": {
                    "br": 1540737,
                    "fid": 0,
                    "size": 40024241,
                    "vd": -45760,
                    "sr": 44100
                },
                "hr": null,
                "a": null,
                "cd": "1",
                "no": 3,
                "rtUrl": null,
                "ftype": 0,
                "rtUrls": [],
                "djId": 0,
                "copyright": 1,
                "s_id": 0,
                "mark": 9007199255011328,
                "originCoverType": 0,
                "originSongSimpleData": null,
                "tagPicList": null,
                "resourceState": true,
                "version": 15,
                "songJumpInfo": null,
                "entertainmentTags": null,
                "awardTags": null,
                "single": 0,
                "noCopyrightRcmd": null,
                "alg": null,
                "displayReason": null,
                "rtype": 0,
                "rurl": null,
                "mst": 9,
                "cp": 7003,
                "mv": 0,
                "publishTime": 1132675200000,
                "videoInfo": {
                    "moreThanOne": false,
                    "video": null
                }
            }
        ],
        "videoIds": null,
        "videos": null,
        "trackIds": [
            {
                "id": 1441524915,
                "v": 6,
                "t": 0,
                "at": 1614999988857,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 1326804162,
                "v": 4,
                "t": 0,
                "at": 1614592023825,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 1326804169,
                "v": 4,
                "t": 0,
                "at": 1613961828597,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 1375171522,
                "v": 9,
                "t": 0,
                "at": 1613716264627,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 1405103306,
                "v": 4,
                "t": 0,
                "at": 1613715497906,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 535709086,
                "v": 7,
                "t": 0,
                "at": 1613715308105,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 1338791509,
                "v": 11,
                "t": 0,
                "at": 1613715223874,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 535709081,
                "v": 9,
                "t": 0,
                "at": 1613715047100,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 22788309,
                "v": 35,
                "t": 0,
                "at": 1613714947901,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 464714353,
                "v": 13,
                "t": 0,
                "at": 1613714772727,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 464714348,
                "v": 13,
                "t": 0,
                "at": 1613714676513,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            },
            {
                "id": 22680695,
                "v": 15,
                "t": 0,
                "at": 1613714433843,
                "alg": null,
                "uid": 757014,
                "rcmdReason": "",
                "sc": null,
                "f": null,
                "sr": null,
                "dpr": null
            }
        ],
        "bannedTrackIds": null,
        "mvResourceInfos": null,
        "shareCount": 1,
        "commentCount": 0,
        "remixVideo": null,
        "newDetailPageRemixVideo": null,
        "sharedUsers": null,
        "historySharedUsers": null,
        "gradeStatus": "NONE",
        "score": null,
        "algTags": null,
        "distributeTags": [],
        "trialMode": 13,
        "displayTags": null,
        "playlistType": "UGC"
    },
    "urls": null,
    "privileges": [
        {
            "id": 1441524915,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 128000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 128000,
            "pc": null,
            "toast": false,
            "flag": 261,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "standard",
            "dlLevel": "none",
            "flLevel": "standard",
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
                    "chargeType": 0
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
        },
        {
            "id": 1326804162,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 128000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 128000,
            "pc": null,
            "toast": false,
            "flag": 261,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "standard",
            "dlLevel": "none",
            "flLevel": "standard",
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
                    "chargeType": 0
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
        },
        {
            "id": 1326804169,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 128000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 128000,
            "pc": null,
            "toast": false,
            "flag": 196869,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "jymaster",
            "playMaxBrLevel": "jymaster",
            "downloadMaxBrLevel": "jymaster",
            "plLevel": "standard",
            "dlLevel": "none",
            "flLevel": "standard",
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
                    "chargeType": 0
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
        },
        {
            "id": 1375171522,
            "fee": 0,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 999000,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 256,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "exhigh",
            "dlLevel": "lossless",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 1405103306,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 128000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 128000,
            "pc": null,
            "toast": false,
            "flag": 260,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "standard",
            "dlLevel": "none",
            "flLevel": "standard",
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
                    "chargeType": 0
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
        },
        {
            "id": 535709086,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 196868,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "jymaster",
            "playMaxBrLevel": "jymaster",
            "downloadMaxBrLevel": "jymaster",
            "plLevel": "exhigh",
            "dlLevel": "none",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 1338791509,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 196868,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "jymaster",
            "playMaxBrLevel": "jymaster",
            "downloadMaxBrLevel": "jymaster",
            "plLevel": "exhigh",
            "dlLevel": "none",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 535709081,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 260,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "exhigh",
            "dlLevel": "none",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 22788309,
            "fee": 1,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 0,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 0,
            "pc": null,
            "toast": false,
            "flag": 197892,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "jymaster",
            "playMaxBrLevel": "jymaster",
            "downloadMaxBrLevel": "jymaster",
            "plLevel": "none",
            "dlLevel": "none",
            "flLevel": "none",
            "rscl": null,
            "freeTrialPrivilege": {
                "resConsumable": true,
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
        },
        {
            "id": 464714353,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 4,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "exhigh",
            "dlLevel": "none",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 464714348,
            "fee": 8,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 320000,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 320000,
            "pc": null,
            "toast": false,
            "flag": 4,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "exhigh",
            "dlLevel": "none",
            "flLevel": "exhigh",
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
                    "chargeType": 0
                },
                {
                    "rate": 192000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
                },
                {
                    "rate": 320000,
                    "chargeUrl": null,
                    "chargeMessage": null,
                    "chargeType": 0
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
        },
        {
            "id": 22680695,
            "fee": 1,
            "payed": 0,
            "realPayed": 0,
            "st": 0,
            "pl": 0,
            "dl": 0,
            "sp": 7,
            "cp": 1,
            "subp": 1,
            "cs": false,
            "maxbr": 999000,
            "fl": 0,
            "pc": null,
            "toast": false,
            "flag": 1284,
            "paidBigBang": false,
            "preSell": false,
            "playMaxbr": 999000,
            "downloadMaxbr": 999000,
            "maxBrLevel": "lossless",
            "playMaxBrLevel": "lossless",
            "downloadMaxBrLevel": "lossless",
            "plLevel": "none",
            "dlLevel": "none",
            "flLevel": "none",
            "rscl": null,
            "freeTrialPrivilege": {
                "resConsumable": true,
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
    "sharedPrivilege": null,
    "resEntrance": null,
    "fromUsers": null,
    "fromUserCount": 0,
    "songFromUsers": null
}`
)

func Test_PlaylistResult(t *testing.T) {
	d := &PlaylistResult{}
	err := fjson.UnmarshalFromString(playlistResult, d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(d)
}

func Test_PlaylistDetail(t *testing.T) {
	playlistURLs.ForEach(func(s string, i int) {
		id, err := PlaylistLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := PlaylistDetail(id, reqData)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}
