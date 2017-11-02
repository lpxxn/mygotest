package urls

import "github.com/mygotest/workspace/webdemo1/src/userinfo"

func init() {

	UrlsGetmap["userinfobyid"] = userinfo.GetUserInfoById
	UrlsGetmap["userinfobyid2/:name"] = userinfo.GetUserInfoById
	UrlsGetmap["callback"] =  userinfo.CallBack
	UrlsGetmap["reqoauth"] = userinfo.RequestOauth
	UrlsGetmap["setusesession"] = userinfo.SetSession
	UrlsGetmap["getusersession"] = userinfo.GetSession

	UrlsPostMap["paginationPersion"] = userinfo.PaginationPerson
}