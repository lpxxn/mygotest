package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
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

	id := bson.NewObjectId()

	runner1, err1 := dbtb1(session, id)
	runner2, err2 := dbtb2(session, id)
	fmt.Println(err1, " err2: ", err2)
	fmt.Println(runner1, runner2)

	if err1 != nil || err2 != nil {
		fmt.Print("have error error info")
		err3 := runner1.Resume(id)
		err4 :=runner2.Resume(id)
		fmt.Println(err3, " err4 ", err4)
	}

}


func dbtb1(session *mgo.Session, id bson.ObjectId) (*txn.Runner, error) {
	db1 := session.DB("lptest")
	tcdemo1 := db1.C("tcdemo")
	runner1 := txn.NewRunner(tcdemo1)
	ops := []txn.Op{{
		C:      "debts",
		Id:     0,
		//Insert: bson.M{"balance": 300},
		Update: bson.M{"$inc": bson.M{"balance": 100}},
	}, {
		C:      "accounts",
		Id:     0,
		Insert: bson.M{"balance": 300},
	}, {
		C:      "accounts",
		Id:     1,
		Insert: bson.M{"balance": 500},
	}, {
		C:      "people",
		Id:     "joe",
		Insert: bson.M{"accounts": []int64{0, 1}},
	}}
	err := runner1.Run(ops, id, nil)
	return runner1, err
}

func dbtb2(session *mgo.Session, id bson.ObjectId) (*txn.Runner ,error) {
	db1 := session.DB("lptest2")
	tcdemo1 := db1.C("tcdemo2")
	runner1 := txn.NewRunner(tcdemo1)
	ops := []txn.Op{{
		C:      "debts",
		Id:     0,
		//Insert: bson.M{"balance": 300},
		Update: bson.M{"$inc": bson.M{"balance": 100}},
	}, {
		C:      "accounts",
		Id:     0,
		Insert: bson.M{"balance": 800},
	}, {
		C:      "accounts",
		Id:     1,
		Insert: bson.M{"balance": 700},
	}, {
		C:      "people",
		Id:     "joe",
		Insert: bson.M{"accounts": []int64{0, 1}},
	},
	{
		C:      "debts",
		Id:     0,
		Update: bson.M{"$inc": bson.M{"balance": "abcd"}},
	},
	}
	err := runner1.Run(ops, id, nil)
	return runner1, err
}
