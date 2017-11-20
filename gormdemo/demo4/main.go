package main

import (
	"fmt"
	"encoding/json"
	"log"
	"github.com/mygotest/gormdemo/utils"
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
		filed1 models.Language
		field2 models.Movie
	}{}
	rdb := db.Table("gotest_languages").Joins("INNER JOIN `gotest_movies` on `gotest_languages`.id = `gotest_movies`.language_id").Select(
		//[]string{"gotest_languages.*", "gotest_movies.*"})//.Scan(&findrev)//.Scan(&findrev)  , "gotest_movies.title"
	[]string{"gotest_languages.id", "gotest_languages.name", "gotest_languages.test_str", "gotest_movies.id", "gotest_movies.title"})//.Find(&findrev)//.Scan(&findrev)
	err := rdb.Error
	fmt.Println(err)

	rows, _ := rdb.Rows()
	for rows.Next() {
		var tempLan models.Language
		var tempMovie models.Movie
		rows.Scan(&tempLan.ID, &tempLan.Name, &tempLan.TestStr, &tempMovie.ID, &tempMovie.LanguageID)
		if tempLan.TestStr == nil {
			fmt.Println("TestStr is null")
		}
		fmt.Println(tempLan)
		fmt.Println(tempMovie)
	}

	jsonVal, err := json.Marshal(findrev)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonVal))
}
