package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	session, err := mgo.Dial("192.168.3.22:27017")
	//session, err := mgo.Dial("rr-mongodb-0001:20000,rr-mongodb-0002:20000,rr-mongodb-0003:20000,rr-mongodb-0004:20000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("mytest").C("products2")
	selecter := bson.M{"no": 1212}

	update := bson.M{
		"$setOnInsert": bson.M{
			"name": "peng",
			"age": 20,
		},
	}

	change := mgo.Change{
		Update: update,
		ReturnNew: true,
		Upsert: true,
	}

	err = coll.Update(selecter, update)
	fmt.Println("insert err:", err)

	doc := bson.M{}
	changeInfo ,err := coll.Find(selecter).Apply(change, &doc)
	if nil != err {
		panic(err)
	}
	fmt.Printf(" changeInfo: %#v \n doc: %#v \n", changeInfo, doc)

	q := coll.Find(nil)//.Select(bson.M{})
	m_data := make([]map[string]interface{}, 0)

	q.All(&m_data)
	fmt.Printf("data len: %d \n data :%#v \n", len(m_data), m_data)

}
