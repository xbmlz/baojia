package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        int            `gorm:"autoIncrement;primary_key" json:"id"`
	Title     string         `json:"title"`
	HTML      string         `gorm:"type:text" json:"html"`
	Text      string         `gorm:"type:text" json:"text"`
	Cover     string         `json:"cover"`
	Author    string         `json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Articles []Article

func CreateArticle(a Article) (err error) {
	return db.Create(&a).Error
}

func GetArticle(id int) (a Article, err error) {
	err = db.First(&a, id).Error
	return
}

func GetArticles() (as Articles, err error) {
	err = db.Order("id desc").Find(&as).Error
	return
}

func UpdateArticle(a Article) (err error) {
	err = db.Save(&a).Error
	return
}

func DeleteArticle(id int) (err error) {
	err = db.Delete(&Article{}, id).Error
	return
}
