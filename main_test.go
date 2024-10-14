package main

import (
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

var reqData *utils.RequestData
var cookie *CookieFile

func init() {
	cookie = NewCookieFile(ncmCookieFile)
	reqData = &utils.RequestData{
		Cookies: cookie.ToHttpCookie(),
	}
}
