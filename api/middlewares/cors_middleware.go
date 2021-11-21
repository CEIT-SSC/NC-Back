package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, content-type, method, x-pingother")
	c.Header("Access-Control-Allow-Credentials", "true")
	if strings.EqualFold("OPTIONS", c.Request.Method) {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}