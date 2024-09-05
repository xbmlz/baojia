package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/api"
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

	r.GET("/admin", api.AdminView)

	apis := r.Group("api")
	{
		apis.GET("product", api.GetProducts)
		apis.POST("price", api.SavePrice)
	}

	r.GET("/", api.IndexView)

	r.Run()
}
