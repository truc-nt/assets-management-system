package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"

	"errors"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerReqParam models.UserRegisterParam
	if err := c.BindJSON(&registerReqParam); err != nil {
		return
	}

	err := services.FindUserByUsername(registerReqParam.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		id, _err := services.CreateUser(registerReqParam)
		if _err != nil {
			h.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, "User created with id: "+strconv.FormatUint(uint64(id), 10))
		return
	}

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusBadRequest, "User already exist")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReqParam models.UserAuthParam
	if err := c.BindJSON(&loginReqParam); err != nil {
		return
	}

	id, role, err := services.AuthenticateUser(loginReqParam)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessLogin(c, id, role)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var id struct{ Id uint32 }
	if err := c.BindJSON(&id); err != nil {
		return
	}

	err := services.SetLogoutUser(id.Id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, id)
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

func (h *AuthHandler) handleSuccessLogin(c *gin.Context, id uint32, role uint32) {
	err := services.SetLoginUser(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, struct {
		Id   uint32
		Role uint32
	}{
		Id:   id,
		Role: role,
	})
}
