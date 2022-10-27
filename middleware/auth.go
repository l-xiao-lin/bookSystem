package middleware

import (
	"bookSystem/pkg"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(200, gin.H{
				"msg": "Auth头部为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bear") {
			c.JSON(200, gin.H{
				"msg": "auth头部中的格式有误",
			})
			c.Abort()
			return
		}
		claims, err := pkg.ParseToken(parts[1])
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
