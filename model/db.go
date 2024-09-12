package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB(dsn string) (err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	return nil
}

func MigrateTable() {
	db.AutoMigrate(&Product{}, &ProductType{}, &Price{}, &User{}, &Sale{}, &Article{})
}
