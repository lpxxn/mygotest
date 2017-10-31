package urls

import "github.com/mygotest/workspace/webdemo1/src/userinfo"

func init() {

	UrlsGetmap["userinfobyid"] = userinfo.GetUserInfoById
	UrlsGetmap["userinfobyid2/:name"] = userinfo.GetUserInfoById
	UrlsGetmap["callback"] = userinfo.CallBack
	UrlsGetmap["reqoauth"] = userinfo.RequestOauth

	UrlsPostMap["paginationPersion"] = userinfo.PaginationPerson
}
