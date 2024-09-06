package main

import (
	"os"

	"github.com/xbmlz/baojia/api"
	"github.com/xbmlz/baojia/model"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("DB_DSN")

	model.InitDB(dsn)
	model.MigrateTable()

	r := api.InitRouter()

	r.Run("0.0.0.0:8080")
}
