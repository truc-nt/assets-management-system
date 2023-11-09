package handlers

import (
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type AssetsHandler struct {
	BaseHandler
}

func NewAssetsHandler() *AssetsHandler {
	return &AssetsHandler{}
}

func (h *AssetsHandler) GetAssets(c *gin.Context) {
	assets, err := services.GetAssets()

	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, assets)
}
