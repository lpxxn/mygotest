package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mygotest/protocoldemo/demo1/protos"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"github.com/gin-contrib/static"
	"runtime"
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


	v1 := r.Group("/v1")
	{
		v1.GET("/user", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"User": "li"})
		})
	}
	// template
	r.LoadHTMLGlob("./public/html/*/*.html")
	//r.Static("/public", "./public")
	//r.Static("/", "./public")
	//r.Use(static.Serve("/", static.LocalFile("./public/html", true)))
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/ie", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{
			"goVersion": "v 1.0",
			"myVersion": runtime.Version(),
		})
		//c.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{"say": "HeHe....."})
	})

	r.GET("/h2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "h2test.html", gin.H{"say": "Hello World"})
	})

	http.ListenAndServe(":9001", r)

}
