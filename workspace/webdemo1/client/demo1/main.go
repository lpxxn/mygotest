package main



import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"github.com/denisenkom/go-mssqldb"

	"time"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "1Qaz2wsx", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "192.168.0.105", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database	  = flag.String("database", "testSplit", "use database")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")

	rows,_ := conn.Query("select * from [User]")
	fmt.Println(rows)

	defer rows.Close()
	for rows.Next() {
		var val1, val2, val3 string
		var val mssql.UniqueIdentifier
		var val4 time.Time
		rows.Scan(&val, &val2, &val4)
		fmt.Println(val.String(), val1, val2, val3, val4)
	}
	fmt.Println("over")
}
