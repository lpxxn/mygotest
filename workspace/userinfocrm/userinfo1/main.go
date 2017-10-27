package main

import (
	"fmt"
	"github.com/mygotest/workspace/userinfocrm/utils"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"time"
	"log"
)

func main() {
	db := utils.SqlDbInstance()
	row, _ := db.Query("select * from MyUser;")
	defer  row.Close()
	fmt.Println(row.Columns())

	var id, name string
	for row.Next() {
		var val1 mysql.NullTime
		row.Scan(&id, &name, &val1)
		fmt.Println(id, name, val1.Time)
	}

	u1 := uuid.NewV4()
	struul := u1.String()
	fmt.Printf("UUIDv4: %s %s\n", u1, struul)

	stmt, err := db.Prepare("INSERT INTO MyUser(id, name, time) VALUE (?, ?, ?)")

	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	randomName := utils.RandStringRunes(5)
	res, err := stmt.Exec(struul, randomName, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id = %d, effected = %d \n", lastId, rowCnt)

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tstmt, err := transaction.Prepare("INSERT INTO MyUser(id, name, time) VALUE (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	tstmt.Exec(uuid.NewV4().String(), "lilililk", time.Now())
	tstmt.Exec(uuid.NewV4().String(), "lilili2", time.Now())
	transaction.Rollback()
	// transaction.Commit()



}
