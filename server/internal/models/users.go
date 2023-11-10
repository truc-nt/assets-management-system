package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Login    bool   `gorm:"column:login" json:"login"`
}

type UserAuthParam struct {
	Username string
	Password string
}

func AuthenticateUser(db *gorm.DB, userAuthParam UserAuthParam) error {
	var user User
	result := db.Model(&User{}).Where(&User{Username: userAuthParam.Username, Password: userAuthParam.Password}).First(&user)
	return result.Error
}

func SetLoginUser(db *gorm.DB, username string) error {
	result := db.Model(&User{}).Where("username = ?", username).Update("login", true)
	return result.Error
}

func SetLogoutUser(db *gorm.DB, username string) error {
	result := db.Model(&User{}).Where("username = ?", username).Update("login", false)
	return result.Error
}
