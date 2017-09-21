package bll

import (
	"io/ioutil"
	"github.com/mygotest/httpdemo1/httptest4/utils"
	"fmt"
	"github.com/mygotest/httpdemo1/httptest4/models"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"google.golang.org/genproto/googleapis/type/date"
)

func GetPrice(jd *models.JdInfo) {
	for index := range jd.Favorite_Products {
		go GetJdPrice(jd.PriceUrl, &jd.Favorite_Products[index])
	}
}

func GetJdPrice(url string, product *models.JdFavoriteProduct) {
	res, err := http.Get(url + product.ProductCode)
	if err == nil {
		msg, _ := ioutil.ReadAll(res.Body)
		var jd models.JdPrice
		json.Unmarshal(msg, &jd)
		t := time.Now()
		d := date.Date{ Year: int32(t.Year()), Month: int32(t.Month()), Day: int32(t.Day())}
		fmt.Printf("京东 商品 %s,当前价格  %s, 期望价格:%g ,编号：%s \n", product.ProductName, jd[0].P, product.FavoritePrice, product.ProductCode)
		if product.SendEmailTime == d && product.SendCount > 0 {
			fmt.Printf("have send %d", product.SendCount)
			return
		}
		if jd != nil && len(jd) > 0  {
			p, _ :=strconv.ParseFloat(jd[0].P, 32)
			pf := float32(p)
			if p > 0 && pf <= product.FavoritePrice {
				fmt.Println("send email")
				go utils.SendEmail(fmt.Sprintf("京东 商品<b> %s </b>,当前价格 <b> %s </b>, 期望价格<b>:%g </b>", product.ProductName, jd[0].P, product.FavoritePrice))
				product.SendCount += 1
				product.SendEmailTime = d
			}
		}
	} else {
		fmt.Errorf("error response")
	}
}