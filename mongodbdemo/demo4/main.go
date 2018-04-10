package main


import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type mydata struct {
	Test bson.Decimal128
}

func (d *mydata) MarshalJSON() {

}

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

	localc := session.DB("wakuangbao2").C("user_coin_day")

	localc.Insert(bson.M{"test": v_deci})
	datas := []mydata{}
	localc.Find(nil).All(&datas)
	fmt.Println(datas)

}
