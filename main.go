package main

import (
	"os"

	"github.com/xbmlz/baojia/api"
	"github.com/xbmlz/baojia/model"
	"github.com/xbmlz/baojia/utils/oss"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("DB_DSN")

	model.InitDB(dsn)
	model.MigrateTable()

	oss.InitMinioClient()

	r := api.InitRouter()

	r.Run("0.0.0.0:8080")
}
