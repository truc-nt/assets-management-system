package handlers

import (
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	BaseHandler
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}

func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	employees, err := services.GetEmployees()
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &employees)
}
