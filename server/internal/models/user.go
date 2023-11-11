package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Login     bool   `gorm:"column:login" json:"login"`
	Role      uint32 `gorm:"column:role" json:"role"`
	Telephone string `gorm:"column:telephone" json:"telephone"`
	DName     string `gorm:"column:dname" json:"dname"`
}

type UserAuthParam struct {
	Username string
	Password string
}

type UserRegisterParam struct {
	Username  string
	Password  string
	Telephone string
	Role      uint32
	DName     string
}

func CreateUser(db *gorm.DB, userRegisterParam UserRegisterParam) (uint32, error) {
	user := User{
		Username:  userRegisterParam.Username,
		Password:  userRegisterParam.Password,
		Role:      userRegisterParam.Role,
		Telephone: userRegisterParam.Telephone,
		Login:     false,
	}
	result := db.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func FindUserByUsername(db *gorm.DB, username string) error {
	var user User
	result := db.Model(&User{}).Where("username = ?", username).First(&user)
	return result.Error
}

func AuthenticateUser(db *gorm.DB, userAuthParam UserAuthParam) (User, error) {
	var user User
	result := db.Model(&User{}).Where(&User{Username: userAuthParam.Username, Password: userAuthParam.Password}).First(&user)
	return user, result.Error
}

func SetLoginUser(db *gorm.DB, id uint32) error {
	result := db.Model(&User{}).Where("id = ?", id).Update("login", true)
	return result.Error
}

func SetLogoutUser(db *gorm.DB, id uint32) error {
	result := db.Model(&User{}).Where("id = ?", id).Update("login", false)
	return result.Error
}
