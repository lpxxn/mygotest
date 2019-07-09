package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := GormPgsqlInit()
	if err != nil {
		panic(err)
	}
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(new(User)).Error; err != nil {
		panic(err)
	}
	u := &User{Name: "💊 adf© 🐶"}
	if err := db.Model(new(User)).Create(u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("user %#v\n", u)
}

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func GormPgsqlInit() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 dbname=lipeng sslmode=disable Timezone=Asia/Shanghai"
	//db, err = sql.Open("mysql", connStr)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
