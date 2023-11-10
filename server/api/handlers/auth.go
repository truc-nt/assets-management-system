package handlers

import (
	"github.com/gin-gonic/gin"
	"server/internal/models"
	"server/internal/services"

	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReqParam models.UserAuthParam
	if err := c.BindJSON(&loginReqParam); err != nil {
		return
	}

	err := services.AuthenticateUser(loginReqParam)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessLogin(c, loginReqParam.Username)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var username struct{ Username string }
	if err := c.BindJSON(&username); err != nil {
		return
	}

	err := services.SetLogoutUser(username.Username)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, username)
}

func (h *AuthHandler) handleError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "User not exist",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
	}
	log.Println(err.Error())
}

func (h *AuthHandler) handleSuccessLogin(c *gin.Context, username string) {
	err := services.SetLoginUser(username)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, username)
}
