package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"autoIncrement;primary_key" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	IsAdmin   bool   `json:"is_admin" default:"false"`
}

func GetUserByUsername(username string) (user *User, exist bool, err error) {
	err = db.Where("username =?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

func CreateUser(user *User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return db.Create(user).Error
}

func GetUserByID(id int) (user *User, err error) {
	err = db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
