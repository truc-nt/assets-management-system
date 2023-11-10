package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	BaseHandler
}

func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	department := models.Department{}

	if err := h.validateInput(c, &department); err != nil {
		return
	}

	if err := services.CreateDepartment(&department); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessCreate(c)
}
