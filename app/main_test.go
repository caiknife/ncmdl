package app

import (
	"github.com/XiaoMengXinX/Music163Api-Go/utils"

	"github.com/caiknife/ncmdl/v2"
)

var reqData *utils.RequestData
var cookie *ncmdl.CookieFile

func init() {
	cookie = ncmdl.NewCookieFile(ncmdl.ncmCookieFile)
	reqData = &utils.RequestData{
		Cookies: cookie.ToHttpCookie(),
	}
}
