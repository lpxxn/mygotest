package favorites

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/mygotest/httpdemo1/httptest4/utils"
)

type JdPrice []struct {
	Op string `json:"op"`
	M  string `json:"m"`
	ID string `json:"id"`
	P  string `json:"p"`
}
const (
	priceUrl = "https://p.3.cn/prices/mgets?skuIds=J_"
)

/*

 */
func GetPrice(product string, myPrice float32) {
	utils.SendEmail()
	//2316993  2316993
	//resp, err := http.Get("https://p.3.cn/prices/mgets?skuIds=J_2316993")
	resp, err := http.Get(priceUrl + product)
	if err == nil {
		msg, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			strBody := string(msg)
			fmt.Println("msg: ", strBody)
			var jd JdPrice
			json.Unmarshal(msg, &jd)
			fmt.Println(jd)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
