package main

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

const (
	br128 = 128000
	br192 = 192000
	br256 = 256000
	br320 = 320000
	br999 = 999000
)

func DownloadSingle(singleID int) (err error) {
	return nil
}

func SingleDetail(singleID int) {

}

// GetSongDownloadURL 获取歌曲下载 URL, 非 VIP 账号可通过此 API 获取部分歌曲的无损音质
func GetSongDownloadURL(data utils.RequestData, id int, bitrate int) (result types.SongDownloadURLData, err error) {
	var options utils.EapiOption
	options.Path = api.SongDownloadURL
	options.Url = "https://music.163.com/eapi/song/enhance/download/url"
	options.Json = api.CreateSongDownloadURLJson(id, bitrate)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
