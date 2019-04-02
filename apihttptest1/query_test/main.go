package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	AdapterMainListURI = "user/order/adaptMainTransList"
	SystemOrderPaidStatus = "Paid"
	SystemOrderRefundingStatus = "Refunding"
	SystemOrderPartialRefundedStatus = "PartialRefunded"
	SystemOrderRefundedStatus = "Refunded"
	SystemOrderRevokedStatus = "Revoked"
)

func main() {
	param := url.Values{}
	param["name"] = []string{"li"}
	strings.NewReader(param.Encode())
	req, _ :=http.NewRequest(http.MethodGet, "http://127.0.0.1:9998/user/orders/adaptMainTransList?status=Refunded&status=Revoked&limit=105", nil)
	query := req.URL.Query()
	query.Set("status", SystemOrderPaidStatus)
	query.Add("status", SystemOrderRefundingStatus)
	query.Add("status", SystemOrderPartialRefundedStatus)
	query.Add("status", SystemOrderRefundedStatus)
	query.Add("status", SystemOrderRevokedStatus)

	query.Set("limit", strconv.Itoa(10))
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	//defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	resp.Body.Close()
}
