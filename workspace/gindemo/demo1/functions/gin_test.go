package funcs_test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// required 需要传入数据，并且不能为默认值
type ReqBindingT struct {
	ID int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// exists 只需要传入的 json 有 字段名就可以  `{"id":0, "name": ""}`
type ReqBindingT2 struct {
	ID *int `json:"id" binding:"exists"`
	Name *string `json:"name" binding:"exists"`
}

type reqBindParam struct {
	ID int `json:"id, omitempty" `
	Name *string `json:"name, omitempty" `
}

func TestBinding(t *testing.T) {
	r := gin.Default()
	r.POST("/b", func(c *gin.Context) {
		p := ReqBindingT{}
		if err := c.BindJSON(&p); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, "ok")
	})

	r.POST("/bc", func(c *gin.Context) {
		p := ReqBindingT2{}
		if err := c.Bind(&p); err != nil {
		//if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, "ok")
	})

	body := ReqBindingT{ID: 0, Name: ""}
	bodyBt, _ := json.Marshal(body)
	// `{"id":0, "name": ""}`
	bodyStr := string(bodyBt)
	fmt.Println(bodyStr)
	// 500 error
	if rspB := postReq("/b", r, bodyStr); rspB.Code != http.StatusOK {
		fmt.Println(rspB.Body.String())
	}

	// 200 success
	if rspC := postReq("/bc", r, bodyStr); rspC.Code != http.StatusOK {
		fmt.Println(rspC.Body.String())
	}

	// 500 error
	//rbody := reqBindParam{}
	//bodyBt, _ = json.Marshal(rbody)
	//bodyStr = string(bodyBt)
	bodyStr = `{"id":0}`
	if rspB := postReq("/bc", r, bodyStr); rspB.Code != http.StatusOK {
		fmt.Println(rspB.Body.String())
	}


}

func postReq(url string, engine *gin.Engine, data string, params  ...func(*http.Request))  *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	for _, p := range params {
		p(req)
	}

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	return w
}
