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

func AdminLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func AppLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(SessionUserKey) == nil {
			c.Redirect(http.StatusFound, "/login.html")
			c.Abort()
		}
		c.Set(CurrentUserKey, session.Get(SessionUserKey))
		c.Next()
	}
}
