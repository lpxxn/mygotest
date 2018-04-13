package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)



func main() {

	mon, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer mon.Close()

	coll := mon.DB("mytest").C("increasecounters")
	//var seq int64 = 1
	//coll.Insert(bson.M{"id": "seqId","seq":seq})

	selecter := bson.M{"id": "seqId"}

	update := bson.M{"$inc": bson.M{"seq": 1}}
	change := mgo.Change{
		Update: update,
		Upsert:true,
		ReturnNew: true,
	}

	doc := bson.M{}
	_, err = coll.Find(selecter).Apply(change, &doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc)

}

func getNextSequence(co *mgo.Collation, name string) {

}