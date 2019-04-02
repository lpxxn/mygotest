package main

import (
	"bytes"
	"fmt"
	"io"
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

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(resp.Body, b)
	// 没有ReadAll 之前 调用 resp.Body.Close()会出错
	resp.Body = ioutil.NopCloser(b)
	tbody, err := ioutil.ReadAll(reader)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tbody))
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	resp.Body.Close()
}
