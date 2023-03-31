package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Check_notMethod(method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != method {
			c.AbortWithStatusJSON(405, gin.H{
				"code": 2001,
				"msg":  "please use " + method + "...",
			})
			//return
		}
	}
}
