package urls

import "github.com/gin-gonic/gin"

var UrlsGetmap map[string]gin.HandlerFunc = make(map[string]gin.HandlerFunc)
var UrlsPostMap map[string]gin.HandlerFunc = make(map[string]gin.HandlerFunc)