package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "li", "123456", "192.168.1.233:3306", "testdb")
	fmt.Println(connStr)
	engine, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic(err)
	}
	engine.SetMaxOpenConns(100)
	engine.SetMaxIdleConns(10)
	engine.ShowSQL(true)

	if err := engine.Sync2(new(User)); err != nil {
		panic(err)
	}
	//u1 := &User{Name: "li", Salt: "asdf", Age:10}
	//u2 := &User{Name: "peng", Salt: "asdf", Age:10}
	//engine.Insert(u1, u2)

	gU := new(User)
	if haveUser, err := engine.Desc("id").Get(gU); err != nil {
		panic(err)
	} else {
		fmt.Println(haveUser)
	}

	DONE := false
	go func() {
		for {
			if DONE {
				return
			}
			//go func() {
			//	newName := RandomStr(5)
			//	newUser := &User{Name: newName, Age: RandomInt(10, 30)}
			//	if _, err := engine.Insert(newUser); err != nil {
			//		panic(err)
			//	}
			//}()
			go func() {
				if err := engine.Sync2(new(User)); err != nil {
					fmt.Println(err)
				}
			}()
			go func() {
				var users []User
				session := engine.NewSession()
				defer session.Rollback()
				if err := session.Desc("id").Find(&users); err != nil {
					fmt.Println(err)
				}
			}()

			if haveUser, err := engine.Desc("id").Get(gU); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(haveUser)
			}
			//<-time.After(time.Millisecond * 500)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch

}

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func (*User) TableName() string {
	return "user"
}

var seed int64

func init() {
	seed = time.Now().UnixNano()
	rand.Seed(seed)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStr(num int) string {
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	b := make([]rune, num)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
