package handlers

import (
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	BaseHandler
}

func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{}
}

func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	departments, err := services.GetDepartments()
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &departments)
}
