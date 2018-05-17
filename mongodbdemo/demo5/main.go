package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	//session, err := mgo.Dial("192.168.3.147:20000")
	session, err := mgo.Dial("rr-mongodb-0001:20000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("urlshortener").C("shorturls")
	// doc := bson.M{"_id": 0, "instock": 1}
	//doc := bson.M{}

	//err = coll.Find(bson.M{"shorturl": "BpLn"}).One(&doc)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(doc)

	datas := make([]bson.M, 0)

	err = coll.Find(bson.M{}).All(&datas)
	fmt.Print(datas)

	for _, v := range datas {
		fmt.Print(v)
	}
//	5.4
// 5.15

}