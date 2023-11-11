package services

import (
	"server/internal/db"
	"server/internal/models"
)

var FindUserByUsername = func(username string) error {
	err := models.FindUserByUsername(db.DB, username)
	return err
}

var CreateUser = func(UserRegisterParam models.UserRegisterParam) (uint32, error) {
	id, err := models.CreateUser(db.DB, UserRegisterParam)
	return id, err
}

var AuthenticateUser = func(userAuthParam models.UserAuthParam) (models.User, error) {
	user, err := models.AuthenticateUser(db.DB, userAuthParam)
	return user, err
}

var SetLoginUser = func(id uint32) error {
	err := models.SetLoginUser(db.DB, id)
	return err
}

var SetLogoutUser = func(id uint32) error {
	err := models.SetLogoutUser(db.DB, id)
	return err
}