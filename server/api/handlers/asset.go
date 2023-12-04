package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

var (
	funcGetAssetHandler = func() services.IAssetService {
		return services.NewAssetService()
	}
)

type GetAssetsParam struct {
	DepartmentId uint32 `form:"department_id"`
}

type AssetHandler struct {
	BaseHandler
	Service services.IAssetService
}

func NewAssetHandler() *AssetHandler {
	return &AssetHandler{
		Service: funcGetAssetHandler(),
	}
}

func (h *AssetHandler) GetAssetsByDepartmentId(c *gin.Context) {
	var param GetAssetsParam
	c.ShouldBindQuery(&param)

	assets, err := h.Service.GetAssetsByDepartmentId(param.DepartmentId)
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

	asset, err := h.Service.GetAssetById(id)
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

	if err := h.Service.CreateAsset(&asset); err != nil {
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

	if err := h.Service.UpdateAsset(id, &asset); err != nil {
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

	if err := h.Service.DeleteAsset(id); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessDelete(c)
}
