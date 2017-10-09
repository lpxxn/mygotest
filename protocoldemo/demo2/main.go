package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mygotest/protocoldemo/demo1/protos"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
)
func main() {
	fmt.Println("test")
	r := gin.Default()
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"a": "hello"})
	})
	r.POST("/testproto1", func(c *gin.Context) {
		//var p1 tutorial.Person
		//if c.Bind(&p1) == nil {
		//	c.JSON(http.StatusOK, gin.H{"p1": p1})
		//}

		var p2 = new(tutorial.Person)
		defer c.Request.Body.Close()
		data, _ := ioutil.ReadAll(c.Request.Body)
		proto.Unmarshal(data, p2)
		fmt.Println(p2);
		p2.Name += "__Test"
		reData, _ := proto.Marshal(p2)
		c.Writer.Write(reData)

	})
	r.POST("/testproto2", func(c *gin.Context) {
		//var p1 tutorial.Person
		//if c.Bind(&p1) == nil {
		//	c.JSON(http.StatusOK, gin.H{"p1": p1})
		//}

		var p2 = new(tutorial.Person)
		defer c.Request.Body.Close()
		data, _ := ioutil.ReadAll(c.Request.Body)
		proto.Unmarshal(data, p2)
		fmt.Println(p2);
		p2.Name += "__Test_JSON_33333"
		//reData, _ := proto.Marshal(p2)
		c.JSON(http.StatusOK, p2)

	})
	http.ListenAndServe(":9000", r)

}
