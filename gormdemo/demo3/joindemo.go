package main

import (
	"github.com/mygotest/gormdemo/utils"
	"fmt"
	//"github.com/mygotest/gormdemo/models"
	"database/sql"
	"github.com/mygotest/gormdemo/models"
)

func main() {
	db := utils.GormInit()
	db.LogMode(true)

	//langs := make([]models.Language, 0)
	//findrev := make([]interface{}, 0)
	// 可以这样查询
	//findrev := []struct {
	//	id string
	//	id2 string
	//	name string
	//	title string
	//}{}
	findrev := [] struct{
		models.Language
		models.Movie
	}{}
	rdb := db.Table("gotest_languages").Joins("INNER JOIN `gotest_movies` on `gotest_languages`.id = `gotest_movies`.language_id").Select(
		[]string{"gotest_languages.id", "gotest_movies.id", "gotest_languages.name", "gotest_movies.title"}).Scan(&findrev)
	err := rdb.Error
	fmt.Println(err)
	fmt.Println(findrev)
	rows, err := rdb.Rows()
	if err != nil {
		panic(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	//rowMaps := make(map[string]interface{})
	fmt.Println("cols : ", cols)
	values := make([]sql.RawBytes, len(cols))

	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scans...)
	}
	fmt.Println(scans)
}