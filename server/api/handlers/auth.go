package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	BaseHandler
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {

}
