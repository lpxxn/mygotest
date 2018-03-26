package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Trb struct {
	Id int
	Name string
	Purchased time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:12qwaszx@tcp(192.168.0.105:3306)/mytest?parseTime=true")
	if err != nil {
		errors.New("open mysql error")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}


	rows, err := db.Query("SELECT * FROM trb3")
	if err!= nil {
		errors.New("query error")
		return
	}

	var datas []Trb
	for rows.Next() {
		var d Trb
		rows.Scan(&d.Id, &d.Name, &d.Purchased)
		fmt.Println(d)
		datas = append(datas, d)
	}
	fmt.Println(datas)

}
