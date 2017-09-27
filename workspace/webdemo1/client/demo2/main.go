package main

import (

	"github.com/mygotest/workspace/webdemo1/src/utils"
	"fmt"

	"github.com/denisenkom/go-mssqldb"
	"time"
)

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

}
