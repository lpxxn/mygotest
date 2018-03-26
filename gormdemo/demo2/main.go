package main

import (
	"fmt"
	"time"

	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var txtDb *gorm.DB = nil

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&loc=Asia%sShanghai", "root", "12qwaszx", "192.168.0.105:3306", "mytest", "%2F")
	fmt.Println(connStr)

	var err error
	txtDb, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic("Open Mysql Error")
	}

	txtDb.DB().SetMaxOpenConns(15000)
	txtDb.DB().SetMaxIdleConns(1000)
	txtDb.DB().SetConnMaxLifetime(time.Duration(time.Minute))
	//txtDb.SingularTable(true)

	txtDb.AutoMigrate(new(MyTools))

	var tools []MyTools

	//toolModel := new(MyTools)
	//rDb := txtDb.Model(toolModel).Find(&tools)
	rDb := txtDb.Find(&tools)
	if rDb.Error != nil {
		panic(rDb.Error)
	}
	fmt.Println(tools)

	var tools2 []MyTools
	rDb = txtDb.Find(&tools2)
	if rDb.Error != nil {
		panic(rDb.Error)
	}
	fmt.Println(tools2)

	var tools3 []MyTools
	rDb = txtDb.Where("name like ?", "%go%").Find(&tools3)
	if rDb.Error != nil {
		panic(rDb.Error)
	}
	fmt.Println(tools3)


	ctime, _ := time.Parse("2006-01-02 15:04:05", "2018-03-11 11:11:11")
	tool1 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82b", Name:"gorm",CreatedA:ctime, CreatedTime: &ctime, CurrTime: time.Now()}, UsedTimes: 3}
	tool2 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82c", Name:"mysql",CreatedA:ctime, CurrTime: time.Now()}, UsedTimes: 100}
	tool3 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82d", Name:"oauth",CreatedA:ctime}, UsedTimes: 100}
	// 如果不传 // CreatedA报错
	//tool4 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82e", Name:"swagger"}, UsedTimes: 100}
	tool4 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82e", Name:"swagger", CreatedA:time.Now()}, UsedTimes: 100}

	txtDb.Create(tool1)
	txtDb.Create(tool2)
	txtDb.Create(tool3)
	rDb = txtDb.Create(tool4)
	if rDb.Error != nil {
		panic(rDb.Error)
	}


	/*
		// delete
		//t1 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82b"}}
		t1 := &MyTools{ModelBase: ModelBase{Id: "02b733a1-940d-4376-8892-2339f783c82c"}}
		rDb = txtDb.Delete(t1)
		if rDb.Error != nil {
			panic(rDb.Error)
		}
		fmt.Println(t1)

	*/

}


type ModelBase struct {
	Id string `gorm:"type:varchar(36);primary_key"`
	Name string `gorm:"type:varchar(100)"`
	CreatedA time.Time `gorm:"type:datetime;"`
	CreatedTime *time.Time `gorm:"type:datetime;" sql:"DEFAULT:NOW()"`
	CurrTime time.Time `gorm:"type:datetime;" sql:"DEFAULT:NOW()"`
}

type MyTools struct {
	ModelBase
	UsedTimes int
}

