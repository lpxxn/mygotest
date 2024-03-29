package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mygotest/workspace/ginswaggerone/models"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/mygotest/workspace/ginswaggerone/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Agm API
// @version 1.0

func main() {
	r := gin.Default()

	r.GET("/lptest/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/lptest/p1/", p1)
	r.GET("/lptest/p2/", p2)

	//r.Run(":8101")
	srv := &http.Server{
		Addr:    ":8101",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go http.ListenAndServe(":8102", r)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("Server is Running....")
	<-quit
	fmt.Println("Shutdown Server....")

}

// @Description get struct array by P1Req
// @ID p1
// @Accept  json
// @Produce  json
// @Param   P1Req  body    models.P1Req     false        " age, name, cur_time:2017-10-11T12:00:12 "
// @Success 200 {object} models.P1Req	"ok"
// @Failure 400 {object} models.APIError "We need ID!!"
// @Failure 404 {object} models.APIError "Can not find ID"
// @Router /lptest/p1/ [post]
func p1(c *gin.Context) {
	var param = &models.P1Req{}
	if c.BindJSON(param) != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}

	fmt.Println(param)
	param.Name = param.Name + "test"
	c.JSON(http.StatusOK, gin.H{"Rev": param})
}

//{
//"age":12134,
//"cur_time": "2001-01-01T10:30:05Z",
//"name":"li"
//}

// @Description get struct array by P1Req
// @ID p2
// @Accept  json
// @Produce  json
// @Param   str1     query    string     false        "param1"
// @Param   str2     query    string     false        "param2"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} models.APIError "We need ID!!"
// @Failure 404 {object} models.APIError "Can not find ID"
// @Router /lptest/p2/ [get]
func p2(c *gin.Context) {
	str1 := c.DefaultQuery("str1", "str11111")
	str2 := c.Query("str2")
	c.JSON(http.StatusOK, gin.H{"str1": str1, "str2": str2})
}
