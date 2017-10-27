package main

import (
	"fmt"
	"github.com/mygotest/workspace/userinfocrm/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
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




}
