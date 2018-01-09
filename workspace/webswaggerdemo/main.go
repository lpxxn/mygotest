package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mygotest/workspace/webswaggerdemo/apis"
	"net/http"
	"os"
	"os/signal"
)

// @APIVersion 0.0.0
// @APITitle title
// @APIDescription description
// @Contact user@domain.com
// @TermsOfServiceUrl http://...
// @License MIT
// @LicenseUrl http://osensource.org/licenses/MIT
func main() {
	r := gin.Default()
	// allow all origins
	r.Use(cors.Default())
	swaggify(r)
	r.GET("/a/b", apis.HandleA)

	go http.ListenAndServe(":5065", r)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("Server is Running....")
	<-quit
	fmt.Println("Shutdown Server .....")

}

// @Title GetStringByInt
// @Description get string by ID
// @Produce  json
// @Success 200 {object} string
// @Router /ab [get]
func a() {

}

///swagger -apiPackage="github.com/mygotest/workspace/webswaggerdemo/apis"  -mainApiFile=github.com/mygotest/workspace/webswaggerdemo/main.go
///delete the basePath in docs.go
