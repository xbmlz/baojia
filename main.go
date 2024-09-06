package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/api"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	model.InitDB(dsn)
	model.MigrateTable()

	r := gin.Default()
	r.Delims("@{", "}")

	r.StaticFS("/public", http.Dir("public"))
	r.LoadHTMLGlob("templates/*")

	store := gormsessions.NewStore(model.GetDB(), true, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login.html", api.AppLoginView)

	adminRouter := r.Group("", middleware.AdminLoginRequired())
	{
		adminRouter.GET("admin", api.AdminView)
	}

	appRouter := r.Group("", middleware.AppLoginRequired())
	{
		appRouter.GET("/", api.AppIndexView)
	}

	apiRouter := r.Group("api")
	{
		apiRouter.GET("product", api.GetProducts)
		apiRouter.POST("price", api.SavePrice)
		apiRouter.POST("register", api.AppRegister)
		apiRouter.POST("login", api.AppLogin)
		apiRouter.POST("logout", api.AppLogout)
	}

	r.Run("0.0.0.0:8080")
}
