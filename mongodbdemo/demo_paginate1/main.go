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

	c := session.DB("wakuangbao").C("user_hash_minute")

	count, err := c.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	query := c.Find(bson.M{})
	cData := []interface{}{}
	query.All(&cData)
	fmt.Println(cData)

	if count > 100 {
		pageSize := 120
		skipValue := 0
		for count > 0 {
			findrev := c.Find(bson.M{}).Skip(skipValue).Limit(pageSize)
			revData := []interface{}{}
			findrev.All(&revData)
			fmt.Println(revData)
			fmt.Println("len: ", len(revData))
			skipValue += pageSize
			count -= pageSize
		}
	}
}
