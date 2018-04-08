package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

const my_time_c = "20060102"

func main() {
	mon, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer mon.Close()

	coll := mon.DB("mytest").C("products")

	selecter := bson.M{
		//"_id": `ObjectId("5aca03e8141b74d92159c257")`,
		"no": 123,
	}

	update := bson.M{
		"$set": bson.M{
			"sku": "abcde",
			"name": "li",
		},
	}

	change := mgo.Change{
		Update: update,
		ReturnNew: true,
		Upsert: true,
	}
	// 没用
	// doc := bson.M{"_id": 0, "instock": 1}
	doc := bson.M{}

	rev, err := coll.Find(selecter).Apply(change, &doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(rev)
	fmt.Println(doc)

	// remove _id
	q := coll.Find(nil).Select(bson.M{"_id": 0})
	m_data := []map[string]interface{}{}
	q.All(&m_data)
	fmt.Println(m_data)


	// have _id
	q2 := coll.Find(nil)
	m_data2 := []map[string]interface{}{}
	q2.All(&m_data2)
	fmt.Println(m_data2)
}
