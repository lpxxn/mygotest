package main

import (
	"fmt"
	"github.com/mygotest/gormdemo/models"
	"github.com/mygotest/gormdemo/utils"
	"github.com/satori/go.uuid"
)

func main() {
	db := utils.GormInit()
	db.SingularTable(true)
	user := &models.User{BaseModel: models.BaseModel{ID: uuid.NewV4().String()}, Name: "li", Age: 18}
	//if has := db.HasTable(&models.User{}); has == false {
	//	db.CreateTable(&models.User{})
	//}
	//er\c
	var err error = db.AutoMigrate(&models.User{}).Error
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var is_blank = db.NewRecord(user)
	fmt.Println(is_blank)
	db.Create(user)
	//var users_data []models.User = make([]models.User, 0)
	var users_data []models.User

	users_data_db := db.Table("gotest_user").Find(&users_data)
	fmt.Println(users_data_db)

}
