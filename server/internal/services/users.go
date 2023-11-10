package services

import (
	"server/internal/db"
	"server/internal/models"
)

var AuthenticateUser = func(userAuthParam models.UserAuthParam) error {
	err := models.AuthenticateUser(db.DB, userAuthParam)
	return err
}

var SetLoginUser = func(username string) error {
	err := models.SetLoginUser(db.DB, username)
	return err
}

var SetLogoutUser = func(username string) error {
	err := models.SetLogoutUser(db.DB, username)
	return err
}
