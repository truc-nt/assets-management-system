package handlers

import (
	"fmt"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type GetAssetsParam struct {
	DepartmentId uint32 `form:"department_id"`
}

type AssetHandler struct {
	BaseHandler
}

func NewAssetHandler() *AssetHandler {
	return &AssetHandler{}
}

func (h *AssetHandler) GetAssetsByEmployeeId(c *gin.Context) {
	var param GetAssetsParam
	c.ShouldBindQuery(&param)

	fmt.Println(param.DepartmentId)

	assets, err := services.GetAssetsByEmployeeId(param.DepartmentId)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &assets)
}

func (h *AssetHandler) GetAssetById(c *gin.Context) {
	id := h.parseId(c, c.Param("asset_id"))
	if id == 0 {
		return
	}
	asset, err := services.GetAssetById(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &asset)
}

func (h *AssetHandler) CreateAsset(c *gin.Context) {
	asset := models.Asset{}

	if err := h.validateInput(c, &asset); err != nil {
		return
	}

	if err := services.CreateAsset(&asset); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessCreate(c)
}

func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	asset := models.Asset{}

	if err := h.validateInput(c, &asset); err != nil {
		return
	}

	id := h.parseId(c, c.Param("asset_id"))
	if id == 0 {
		return
	}

	if err := services.UpdateAsset(id, &asset); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessUpdate(c)
}

func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	id := h.parseId(c, c.Param("asset_id"))
	if id == 0 {
		return
	}

	if err := services.DeleteAsset(id); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessDelete(c)
}
