package userinfo

import (

	//"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	Name string
	Age int
}

// Get User
func GetUserInfoById(c *gin.Context) {
	strpar := c.DefaultQuery("id", "1")
	// param in path
	strname := c.Param("name")
	id, _ := strconv.Atoi(strpar)
	users := []User{User{Name:"li", Age:10}, {Name:strname, Age: 8}}
	if id > 2 {
		users = append(users, User{Name: "Na", Age: 10})
	}
	c.JSON(200, users)
	//rvjson, _ := json.Marshal(users)
	//c.JSON(200, string(rvjson))
}
