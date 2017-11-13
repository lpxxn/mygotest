package main

import (
	"github.com/mygotest/gormdemo/models"
	"github.com/mygotest/gormdemo/utils"
	"github.com/satori/go.uuid"
)

func main() {
	db := utils.SqldbInit()
	user := &models.User{BaseModel: models.BaseModel{ID: uuid.NewV4().String()}, Name: "li"}
	db.AutoMigrate(&models.User{})
	//db.c

}
