package urls

import "github.com/mygotest/workspace/webdemo1/src/ordersinfo"

func init() {
	UrlsPostMap["orderbyid"] = ordersinfo.OrdersById
}