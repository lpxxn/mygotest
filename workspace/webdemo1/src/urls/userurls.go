package urls

import "github.com/mygotest/workspace/webdemo1/src/userinfo"

func init() {

	UrlsGetmap["userinfobyid"] = userinfo.GetUserInfoById
}