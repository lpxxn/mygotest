package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

func Serve(urlPrefix string, fs packr.Box) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.Has(c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
