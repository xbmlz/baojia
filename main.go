package main

import (
	"log"
	"os"

	"github.com/xbmlz/baojia/api"
	"github.com/xbmlz/baojia/cron"
	"github.com/xbmlz/baojia/model"
	"github.com/xbmlz/baojia/pkg/oss"
	"github.com/xbmlz/baojia/pkg/wechat"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("DB_DSN")

	if err := model.InitDB(dsn); err != nil {
		log.Fatal(err)
	}

	model.MigrateTable()

	oss.InitMinioClient()

	go func() {
		wechat.InitWeChatBot()
	}()

	// cron
	go func() {
		cron.Run()
	}()

	r := api.InitRouter()

	log.Fatal(r.Run("0.0.0.0:8080"))
}
