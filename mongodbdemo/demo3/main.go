package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"math"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func main() {

	//session, err := mgo.Dial("192.168.3.147:20000")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var input_num float64 = 10.3798
	fv := strconv.FormatFloat(input_num, 'f', -1, 64)
	//fv := strconv.FormatFloat(input_num, 'f', 2, 64)
	v_deci, err := bson.ParseDecimal128(fv)
	fmt.Println(v_deci, "  err: ", err, "  detail: ")
	rv_float,_ := strconv.ParseFloat(v_deci.String(), 64)
	fmt.Println(rv_float)

	localc := session.DB("wakuangbao").C("user_coin_day")

	localc.Insert(bson.M{"round": round(1.2345, 2) + 123})
	localc.Insert(bson.M{"round": float64(round(5.2345, 3) + 555.65)})


	localc.Insert(bson.M{"toFixed": round(1.2345, 2)})
	localc.Insert(bson.M{"toFixed": round(5.2345, 3)})
	err = localc.Insert(bson.M{"toFixed": v_deci})
	if err != nil {
		fmt.Println(err)
	}


	data := []map[string]interface{}{}
	localc.Find(nil).All(&data)
	fmt.Println(data)
}

func round(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10n)*pow10n) / pow10n
}


func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round2(num * output)) / output
}

func round2(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
