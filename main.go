package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/api"
	"github.com/xbmlz/baojia/model"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed templates/* public/*
var f embed.FS

func main() {
	dsn := os.Getenv("DB_DSN")
	model.InitDB(dsn)
	model.MigrateTable()

	r := gin.Default()

	// 自定义模板函数
	funcMap := template.FuncMap{}

	// embed files
	tmpl := template.New("").Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseFS(f, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)
	r.Delims("@{", "}")

	fp, _ := fs.Sub(f, "public")
	r.StaticFS("/public", http.FS(fp))
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
