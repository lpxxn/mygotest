package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"fmt"
	"flag"
	"time"
)

var (
	server = flag.String("server", "54.223.243.21:3306", "db ip")
	database = flag.String("database", "hydra", "user database")
	user = flag.String("user", "mtk_backoffice", "db users")
	password = flag.String("password", "q2Ayy2qPE6wKyqan", "the database password")
)

func init() {
	SqlDbInstance()
}

var db *sql.DB = nil
var once sync.Once

func SqlDbInstance() *sql.DB {
	once.Do(func() {
		connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", *user, *password, *server, *database)
		fmt.Println("connStr: ", connectStr)
		var err error
		db, err = sql.Open("mysql", connectStr)

		if err != nil  {
			panic("Open connect filed")
		}

		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)
		db.SetConnMaxLifetime(time.Minute)
	})
	return db
}