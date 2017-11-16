package main

import (
	"fmt"
	"github.com/mygotest/gormdemo/models"
	"github.com/mygotest/gormdemo/utils"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func main() {
	db := utils.GormInit()
	db.AutoMigrate(new(models.Language), new(models.Movie), new(models.Artist))
	initData()
}

func initData() {
	db := utils.GormInit()

	langualg1 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now()}, Name: "tamil"}
	langualg2 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now()}, Name: "hindi"}
	langualg3 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now()}, Name: "english"}
	languages := []models.Language{langualg1, langualg2, langualg3}
	vals := []interface{}{}
	//sqlStr := "REPLACE INTO gotest_languages (`id`, `name`, `created_at`, `update_time`, `test_time`) VALUES"
	sqlStr := "REPLACE INTO gotest_languages (`id`, `name`, `created_at`, `update_time`) VALUES"
	for _, lang := range languages {
		val := []interface{}{
			lang.ID,
			lang.Name,
			lang.CreatedAt,
			//nil,
			time.Now(),
		}
		q := strings.Repeat("?,", len(val))
		q = q[0 : len(q)-1]
		sqlStr += fmt.Sprintf("(%s),", q)
		vals = append(vals, val...)
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	fmt.Println(sqlStr)
	fmt.Println(vals)
	fmt.Println(vals...)
	stmt, err := db.CommonDB().Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	rel, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}
	rows, count := rel.RowsAffected()
	fmt.Println(rows, count)

}
