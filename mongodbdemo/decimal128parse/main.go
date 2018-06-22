package main

import (
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"encoding/json"
)

type Data struct {
 	Amount MyDecimal
 	Name string
}



type MyDecimal struct {
	float64
}

func (i *MyDecimal) GetBSON() (interface{}, error) {
	svr := strconv.FormatFloat(i.float64, 'f', -1, 64)
	return bson.ParseDecimal128(svr)
}

func (i *MyDecimal) SetBSON(raw bson.Raw) error {
	var v bson.Decimal128
	err :=raw.Unmarshal(&v)
	fmt.Println("set  err", v, "  ", v.String())
	if err == nil {
		fv, _ :=strconv.ParseFloat(v.String(), 64)
		i.float64  = fv
	}
	fmt.Println("vvv", i)
	return err
}

func (i *MyDecimal) UnmarshalJSON(d []byte) error {
	var v float64
	fmt.Println("unmarshaljson, ", i.float64)
	err := json.Unmarshal(d, &v)
	if err != nil {
		i.float64 = v
	}

	return err
}

func (i *MyDecimal) MarshalJSON()([]byte, error) {
	fmt.Println("------", i.float64, "  ", i)

	return json.Marshal(i.float64)
}
