package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

var db *gorm.DB = nil
var sqlonce sync.Once

func SqldbInit() *gorm.DB {
	sqlonce.Do(func() {

		//connStr := fmt.Sprintf("server=%s;password=%s;port=%d;database=%s;", "192.168.0.105", "12qwaszx", 3306, "GoOrmTest")
		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "12qwaszx", "192.168.0.105:3306", "GoOrmTest")
		fmt.Printf(connStr)

		var err error
		//db, err = sql.Open("mysql", connStr)
		db, err = gorm.Open("mysql", connStr)
		if err != nil {
			panic("Open mssql error")
		}
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "gotest_" + defaultTableName
		}
		db.DB().SetMaxOpenConns(15000)
		db.DB().SetMaxIdleConns(1000)
		db.DB().SetConnMaxLifetime(time.Duration(time.Minute))
	})
	return db
}
