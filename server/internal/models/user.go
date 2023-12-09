package models

import (
	"fmt"

	"gorm.io/gorm"
)

type GetUsersParam struct {
	Role uint32 `form:"role"`
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

type User struct {
	ID        uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Login     bool   `gorm:"column:login" json:"login"`
	Role      uint32 `gorm:"column:role" json:"role"`
	Telephone string `gorm:"column:telephone" json:"telephone"`
	DName     string `gorm:"column:dname" json:"dname"`
}

type IUserRepository interface {
	AuthenticateUser(db *gorm.DB, userAuthParam UserAuthParam) (User, error)
	SetLoginUser(id uint32) error
	SetLogoutUser(id uint32) error
	GetUsers(param *GetUsersParam) ([]*User, error)
	GetUserById(id uint32) (*User, error)
	CreateUser(userRegisterParam UserRegisterParam) (uint32, error)
	UpdateUser(user *User) error
	FindUserByUsername(username string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(userRegisterParam UserRegisterParam) (uint32, error) {
	user := User{
		Username:  userRegisterParam.Username,
		Password:  userRegisterParam.Password,
		Role:      userRegisterParam.Role,
		Telephone: userRegisterParam.Telephone,
		DName:     userRegisterParam.DName,
		Login:     false,
	}
	result := r.DB.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func (r *UserRepository) FindUserByUsername(username string) error {
	var user User
	result := r.DB.Model(&User{}).Where("username = ?", username).First(&user)
	return result.Error
}

func (r *UserRepository) AuthenticateUser(db *gorm.DB, userAuthParam UserAuthParam) (User, error) {
	var user User
	result := db.Model(&User{}).Where(&User{Username: userAuthParam.Username, Password: userAuthParam.Password}).First(&user)
	return user, result.Error
}

func (r *UserRepository) SetLoginUser(id uint32) error {
	result := r.DB.Model(&User{}).Where("id = ?", id).Update("login", true)
	return result.Error
}

func (r *UserRepository) SetLogoutUser(id uint32) error {
	result := r.DB.Model(&User{}).Where("id = ?", id).Update("login", false)
	return result.Error
}

func (r *UserRepository) GetUsers(param *GetUsersParam) ([]*User, error) {
	var users []*User
	if param.Role == 0 {
		err := r.DB.Model(&User{}).Select("id", "username", "dname").Find(&users).Error
		return users, err
	}
	err := r.DB.Where(fmt.Sprintf("role = %d", param.Role)).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserById(id uint32) (*User, error) {
	var user User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) UpdateUser(user *User) error {
	return r.DB.Updates(&user).Error
}
