package main

import (
	"encoding/json"
	"fmt"
)

const revData = `
{
    "resultCode": "OK",
    "resultDescription": "SUCCESS",
    "data": {
        "tradeNo": 114993140347196113,
        "next": {
            "id": 114993140347196113,
            "createTime": 1551770052,
            "amountInCent": 3,
            "status": "New",
            "orderId": "MAIN-Recharge-155211278720180227",
            "clientId": 237879989763099715,
            "paymentServiceId": 223381299346328643,
            "effectiveAmountInCent": ,
            "effectiveTime": 0,
            "finishTime": 0,
            "subject": "在线支付",
            "description": "在线支付",
            "userIdentifier": "2dd22fedb832"
        }
    }
}
`

type A struct {
	ResultCode string `json:"resultCode"`
	Data       struct {
		TradeNo int64 `json:"tradeNo"`
	} `json:"data"`
}

func main() {
	// interface 里的ID是不正确的
	mapData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(revData), &mapData); err != nil {
		panic(err)
	}
	a := A{}
	if err := json.Unmarshal([]byte(revData), &a); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", mapData)
	revStr, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(revStr))
}
