package jwt

import (
	"example.com/blog/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var message string
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = 100
			message = "token为空"
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 401
				message = "token无效"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
				message = "token已过期"
			}
		}
		if code != 0 {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  message,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
