package userinfo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mygotest/workspace/webdemo2/server/parammodes"
	"github.com/mygotest/workspace/webdemo2/utils"
	"net/http"
)

func SetUserInfo(c *gin.Context) {
	var param = &parammodes.UserParam{}
	if c.BindJSON(param) != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}
	pjson, _ := json.Marshal(param)
	str, err := utils.Cluster.Set("gouser", pjson, 0).Result()
	fmt.Println(str, err)
	c.JSON(http.StatusOK, gin.H{"Status": true})
}

func GetUserInfo(c *gin.Context) {
	str, err := utils.Cluster.Get("gouser").Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}
	u := &parammodes.UserParam{}
	json.Unmarshal([]byte(str), u)
	c.JSON(http.StatusOK, gin.H{"Status": true, "Data": u})

}
