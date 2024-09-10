package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/utils/token"
)

const (
	SessionUserKey = "user"
	CurrentUserKey = "current_user"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(SessionUserKey) == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set(CurrentUserKey, session.Get(SessionUserKey))
		c.Next()
	}
}

func JwtAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set(CurrentUserKey, claims)
		c.Next()
	}
}
