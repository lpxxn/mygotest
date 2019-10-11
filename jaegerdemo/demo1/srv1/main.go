package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygotest/jaegerdemo/demo1/common"
	"github.com/mygotest/jaegerdemo/demo1/middleware"
	"github.com/mygotest/jaegerdemo/demo1/utils"
	"github.com/opentracing/opentracing-go"
)

func main() {
	gin.SetMode(gin.DebugMode)
	tracer, closer := utils.NewJaegerTracer("srv1", common.JaegerHostPort)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	engine := gin.New()
	engine.Use(middleware.SetupRootTrace())

	engine.GET("userInfo", UserInfo)
	if err := engine.Run(":9654"); err != nil {
		log.Fatal(err)
	}
}

func UserInfo(c *gin.Context) {
	revUserInfo := common.UserInfo{}
	utils.RandomSleep(0, 1)

	_, err := utils.GetRequestAndUnmarshalResponseData("http://127.0.0.1:9655/remoteUserInfo", &revUserInfo, func(r *http.Request) {
		span, found := middleware.GetSpan(c)
		if found {
			utils.InjectTraceID(span.Context(), r.Header)
		}
	})
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"Err": err})
		return
	}
	utils.RandomSleep(0, 2)
	if _, err := utils.GetRequest("http://127.0.0.1:9656/remote3", func(r *http.Request) {
		span, found := middleware.GetSpan(c)
		if found {
			utils.InjectTraceID(span.Context(), r.Header)
		}
	}); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"Err": err})
		return
	}
	c.JSON(http.StatusOK, revUserInfo)
}
