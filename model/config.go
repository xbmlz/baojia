package model

import "time"

type Config struct {
	ID        int       `gorm:"autoIncrement;primary_key" json:"id"`
	ConfigKey string    `json:"config_key"`
	ConfigVal string    `json:"config_val"`
	Status    int       `json:"status"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Configs []Config

func GetConfigByKey(key string) (Config, error) {
	var config Config
	err := db.Where("config_key = ?", key).First(&config).Error
	return config, err
}

func GetConfigsByKey(key string) (Configs, error) {
	var configs Configs
	err := db.Where("config_key = ?", key).Find(&configs).Error
	return configs, err
}
