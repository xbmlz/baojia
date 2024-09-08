package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		}
		c.Set(CurrentUserKey, session.Get(SessionUserKey))
		c.Next()
	}
}
