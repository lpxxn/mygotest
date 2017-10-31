package bll

import (
	"encoding/json"
	"fmt"
	"github.com/mygotest/httpdemo1/httptest4/models"
	"github.com/mygotest/httpdemo1/httptest4/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetPrice(jd *models.JdInfo) {
	for index := range jd.Favorite_Products {
		go GetJdPrice(&jd.Favorite_Products[index])
	}
}

func RandomUpdatePduid() {
	config := utils.AppConfigInstance()
	str := config.JdProductInfo.Pduid
	index := utils.Random(1, len(str))
	value := utils.Random(1, 9)
	config.JdProductInfo.Pduid = utils.ReplaceAtIndex(str, strconv.Itoa(value), index)
}

func GetJdPrice(product *models.JdFavoriteProduct) {
	config := utils.AppConfigInstance()
	url := fmt.Sprintf(config.JdProductInfo.PriceUrl, config.JdProductInfo.Pduid)
	fmt.Println("url", url)
	res, err := http.Get(url + product.ProductCode)
	if err == nil {
		msg, _ := ioutil.ReadAll(res.Body)
		var jd models.JdPrice
		json.Unmarshal(msg, &jd)
		if jd == nil || len(jd) == 0 {
			fmt.Println("京东接口出现问题，没有返回数据")
			RandomUpdatePduid()
			return
		}
		t := time.Now()
		fmt.Printf("京东 商品 %s,当前价格  %s, 期望价格:%g ,编号：%s \n", product.ProductName, jd[0].P, product.FavoritePrice, product.ProductCode)
		if product.SendCount > 0 {
			fmt.Printf("have send %d \n", product.SendCount)
			return
		}

		p, _ := strconv.ParseFloat(jd[0].P, 32)
		pf := float32(p)
		if p > 0 && pf <= product.FavoritePrice {
			strchan := make(chan string)
			fmt.Println("send email")
			go utils.SendEmail(strchan)
			strchan <- fmt.Sprintf("京东 商品<b> %s </b>,当前价格 <b> %s </b>, 期望价格<b>:%g </b>, 编号：%s ", product.ProductName, jd[0].P, product.FavoritePrice, product.ProductCode)
			product.SendCount += 1
			product.SendEmailTime = t
		}

	} else {
		fmt.Errorf("error response")
	}
}
