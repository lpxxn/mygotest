package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Comunities_coin_day struct {
	Cid int64
	Day string
	SsCoin float64
	SsZuan float64
	YsCoin float64
	YsZuan float64
	//Day time.Time
}
func main() {
	session, err := mgo.Dial("192.168.3.22:20000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("wakuangbao").C("comunities_coin_day")

	query := c.Find(bson.M{})
	cData := []Comunities_coin_day{}
	query.All(&cData)
	fmt.Println(cData)


	localMo, err := mgo.Dial("127.0.0.1:20000")
	if err != nil {
		panic(err)
	}
	defer localMo.Close()

	coll := localMo.DB("mydb").C("mycoll")

	//err = coll.Insert(bson.M{"n": 1}, bson.M{"n": 2}, bson.M{"n": 3})


	bulk := coll.Bulk()
	bulk.Upsert(bson.M{"n": 20}, bson.M{"$set": bson.M{"n": 22}})
	bulk.Upsert(bson.M{"n": 40}, bson.M{"$set": bson.M{"n": 45}}, bson.M{"n": 3}, bson.M{"$set": bson.M{"n": 30}})
	r, err := bulk.Run()

	fmt.Println(r, err)

	localc := localMo.DB("wakuangbao").C("comunities_coin_day")
	localBulk := localc.Bulk()
	//for i, dLen := 0, len(cData); i < dLen; i++ {
	//	//localBulk.Upsert(cData[i], bson.M{"$set": cData[i]})
	//	localBulk.Upsert(bson.M{"cid": cData[i].Cid, "day": cData[i].Day}, bson.M{"$set": cData[i]})
	//}

	for _, v := range cData{
		//localBulk.Upsert(cData[i], bson.M{"$set": cData[i]})
		localBulk.Upsert(bson.M{"cid": v.Cid, "day": v.Day}, bson.M{"$set": v})
	}
	//localBulk.Upsert(cData)
	r, err = localBulk.Run()
	fmt.Println(r, err)

	//b := c.Bulk()
	//b.Insert()
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)
}

