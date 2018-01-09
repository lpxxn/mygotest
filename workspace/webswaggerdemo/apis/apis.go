// @SubApi Test API [/a]
package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Title GetStruct3
// @Description get struct3
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /a/b [get]
func HandleA(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"p1": "hey"})
}
