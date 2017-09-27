package utils

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"flag"
	"fmt"
	"time"
	"sync"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "1Qaz2wsx", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "192.168.0.105", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database	  = flag.String("database", "testSplit", "use database")
)

func init() {
	SqlDbInstance()
}


var db *sql.DB = nil
var once sync.Once
func SqlDbInstance() *sql.DB {
	once.Do(func() {
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *server, *user, *password, *port, *database)
		if *debug {
			fmt.Printf(" connString:%s\n", connString)
		}
		var err error
		db, err = sql.Open("mssql", connString)
		if err != nil {
			panic("Open connection failed")
		}
		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)
		db.SetConnMaxLifetime(time.Duration(time.Minute))
	})
	return db
}
