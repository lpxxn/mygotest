package main

import (
	"fmt"
	"github.com/mygotest/workspace/webdemo1/src/utils"

	"github.com/denisenkom/go-mssqldb"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	db := utils.SqlDbInstance()
	//err := db.Ping()
	//fmt.Println(err.Error())
	row, _ := db.Query("select * from [User]")
	defer row.Close()
	fmt.Println(row.Columns())
	var id mssql.UniqueIdentifier
	var name string
	for row.Next() {
		var val1 time.Time
		row.Scan(&id, &name, &val1)
		fmt.Println(id.String(), name, val1)
	}

	txn, err := db.Begin()
	if err != nil {
		panic("error stmt")
	}

	stmt, err := txn.Prepare(mssql.CopyIn("[User]", mssql.MssqlBulkOptions{}, "Name", "CTime"))

	total := rand.Intn(50)
	for ; total > 0; total-- {
		nameLen := rand.Intn(10)
		stmt.Exec(generateString(nameLen), time.Now())
		time.Sleep(time.Duration(time.Millisecond))
	}
	result, err := stmt.Exec()
	if err != nil {
		panic(err)
	}

	err = stmt.Close()
	if err != nil {
		panic(err)
	}

	err = txn.Commit()
	if err != nil {
		panic(err)
	}
	rowCount, _ := result.RowsAffected()
	log.Printf("%d row copied\n", rowCount)
	//db.Exec("insert into [User] (Name, CTime) values (?1, ?2)", "peng", time.Now())
}

func generateString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterLen := len(letters)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(letterLen)]
	}
	return string(b)
}
