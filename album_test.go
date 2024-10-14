package main

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/samber/lo"
)

const (
	albumResult = `{
    "info": {
        "resourceType": 3,
        "commentCount": 0,
        "likedCount": 0,
        "shareCount": 3,
        "threadId": "R_AL_3_123837273"
    },
    "songs": [
        {
            "name": "Shivers",
            "id": 1824625653,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 209160,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 8368631,
                "vd": 270,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 5021196,
                "vd": 2896,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 3347479,
                "vd": 4683,
                "sr": 44100
            },
            "sq": {
                "br": 1556050,
                "fid": 0,
                "size": 40682979,
                "vd": 272,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Surrender Dear",
            "id": 1824625065,
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
            "pop": 5,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 254303,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 10175260,
                "vd": -1303,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 6105173,
                "vd": 1338,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 4070130,
                "vd": 3147,
                "sr": 44100
            },
            "sq": {
                "br": 1403688,
                "fid": 0,
                "size": 44620319,
                "vd": -1304,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 2,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Soft Winds",
            "id": 1824625654,
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
            "pop": 5,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 218642,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 8747929,
                "vd": 634,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 5248775,
                "vd": 3270,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 3499198,
                "vd": 5089,
                "sr": 44100
            },
            "sq": {
                "br": 1502700,
                "fid": 0,
                "size": 41069348,
                "vd": 638,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 3,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
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
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
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
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Six Appeal",
            "id": 1824625067,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 225565,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 9024827,
                "vd": 480,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 5414914,
                "vd": 3111,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 3609957,
                "vd": 4873,
                "sr": 44100
            },
            "sq": {
                "br": 1491405,
                "fid": 0,
                "size": 42051272,
                "vd": 482,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 5,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Sunday",
            "id": 1824625655,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 185679,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 7430313,
                "vd": 393,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 4458205,
                "vd": 3030,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 2972151,
                "vd": 4828,
                "sr": 44100
            },
            "sq": {
                "br": 1484360,
                "fid": 0,
                "size": 34451897,
                "vd": 391,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "El Salon de Gutbucket",
            "id": 1824625068,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 229453,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 9180517,
                "vd": 266,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 5508328,
                "vd": 2897,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 3672233,
                "vd": 4682,
                "sr": 44100
            },
            "sq": {
                "br": 1550446,
                "fid": 0,
                "size": 44469390,
                "vd": 266,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "On the Sunny Side of the Street",
            "id": 1824625656,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 244988,
            "h": {
                "br": 320001,
                "fid": 0,
                "size": 9802231,
                "vd": 949,
                "sr": 44100
            },
            "m": {
                "br": 192001,
                "fid": 0,
                "size": 5881356,
                "vd": 3582,
                "sr": 44100
            },
            "l": {
                "br": 128001,
                "fid": 0,
                "size": 3920919,
                "vd": 5380,
                "sr": 44100
            },
            "sq": {
                "br": 1511284,
                "fid": 0,
                "size": 46281004,
                "vd": 942,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 8,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Jacquet in the Box",
            "id": 1824625069,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 187395,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 7498231,
                "vd": 400,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 4498956,
                "vd": 3036,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 2999319,
                "vd": 4836,
                "sr": 44100
            },
            "sq": {
                "br": 1539100,
                "fid": 0,
                "size": 36052542,
                "vd": 404,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 9,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "If I Had You",
            "id": 1824625070,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 5,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 268600,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 10746819,
                "vd": 566,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 6448109,
                "vd": 3207,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 4298754,
                "vd": 5000,
                "sr": 44100
            },
            "sq": {
                "br": 1465971,
                "fid": 0,
                "size": 49220111,
                "vd": 568,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 10,
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
            "version": 5,
            "songJumpInfo": null,
            "entertainmentTags": null,
            "awardTags": null,
            "single": 0,
            "noCopyrightRcmd": null,
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Coquette",
            "id": 1824625657,
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
            "pop": 5,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 6,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 171395,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 6858754,
                "vd": 1139,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 4115270,
                "vd": 3778,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 2743528,
                "vd": 5578,
                "sr": 44100
            },
            "sq": {
                "br": 1515445,
                "fid": 0,
                "size": 32467554,
                "vd": 894,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 11,
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
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        },
        {
            "name": "Esquire Jump",
            "id": 1824625071,
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
            "pop": 10,
            "st": 0,
            "rt": "",
            "fee": 1,
            "v": 6,
            "crbt": null,
            "cf": "",
            "al": {
                "id": 123837273,
                "name": "Groove at First Sight",
                "picUrl": "https://p1.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
                "tns": [],
                "pic_str": "109951165776330548",
                "pic": 109951165776330540
            },
            "dt": 146840,
            "h": {
                "br": 320000,
                "fid": 0,
                "size": 5876550,
                "vd": 838,
                "sr": 44100
            },
            "m": {
                "br": 192000,
                "fid": 0,
                "size": 3525947,
                "vd": 3448,
                "sr": 44100
            },
            "l": {
                "br": 128000,
                "fid": 0,
                "size": 2350646,
                "vd": 5221,
                "sr": 44100
            },
            "sq": {
                "br": 1585153,
                "fid": 0,
                "size": 29095675,
                "vd": 845,
                "sr": 44100
            },
            "hr": null,
            "a": null,
            "cd": "01",
            "no": 12,
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
            "rtype": 0,
            "rurl": null,
            "mst": 9,
            "cp": 1416336,
            "mv": 0,
            "publishTime": 1615478400000,
            "videoInfo": {
                "moreThanOne": false,
                "video": null
            }
        }
    ],
    "preSellSongIds": [],
    "code": 200,
    "album": {
        "name": "Groove at First Sight",
        "id": 123837273,
        "type": "专辑",
        "size": 12,
        "picId": 109951165776330540,
        "blurPicUrl": "https://p2.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
        "companyId": 0,
        "pic": 109951165776330540,
        "picUrl": "https://p2.music.126.net/PODxUEKyHcqCkE-oIHyXOw==/109951165776330548.jpg",
        "publishTime": 1615478400000,
        "description": "",
        "tags": "",
        "company": "Snibor Records",
        "briefDesc": "",
        "artist": {
            "name": "Enric Peidro & Jonathan Stout",
            "id": 47024056,
            "picId": 109951167739299220,
            "img1v1Id": 0,
            "briefDesc": "",
            "picUrl": "https://p2.music.126.net/WZNNpBUxMgIuaSDAxDIWCQ==/109951167739299217.jpg",
            "img1v1Url": "https://p2.music.126.net/6y-UleORITEDbvrOLV0Q8A==/5639395138885805.jpg",
            "albumSize": 1,
            "alias": [],
            "trans": "",
            "musicSize": 12,
            "topicPerson": 0,
            "picId_str": "109951167739299217"
        },
        "songs": [],
        "alias": [],
        "status": 1,
        "copyrightId": 1416336,
        "commentThreadId": "R_AL_3_123837273",
        "artists": [
            {
                "name": "Enric Peidro & Jonathan Stout",
                "id": 47024056,
                "picId": 0,
                "img1v1Id": 0,
                "briefDesc": "",
                "picUrl": "",
                "img1v1Url": "https://p2.music.126.net/6y-UleORITEDbvrOLV0Q8A==/5639395138885805.jpg",
                "albumSize": 0,
                "alias": [],
                "trans": "",
                "musicSize": 0,
                "topicPerson": 0
            }
        ],
        "subType": "录音室版",
        "transName": null,
        "onSale": false,
        "mark": 8192,
        "gapless": 0,
        "dolbyMark": 0,
        "picId_str": "109951165776330548",
        "locked": false
    }
}`
)

func Test_AlbumResult(t *testing.T) {
	d := &AlbumResult{}
	err := fjson.UnmarshalFromString(albumResult, d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(d.Album)
}

func Test_AlbumDetail_Random_Order(t *testing.T) {
	// const album = "https://music.163.com/album?id=153840942&userid=757014"
	const album = "https://music.163.com/album?id=195093556&userid=757014"
	id, err := AlbumLinkID(album)
	if err != nil {
		t.Error(err)
		return
	}

	detail, err := AlbumDetail(id, reqData)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(detail)

	songIDs := lo.Map[*SingleInfo, int](detail, func(item *SingleInfo, index int) int {
		return item.ID
	})
	t.Log(songIDs)
	info, err := DownloadLink(songIDs, reqData)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func Test_AlbumDetail(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		id, err := AlbumLinkID(s)
		if err != nil {
			t.Error(err)
			return
		}
		detail, err := AlbumDetail(id, reqData)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(detail)
	})
}
