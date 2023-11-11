package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadDepartmentRoute(route *gin.RouterGroup) {
	departmentHandler := handlers.NewDepartmentHandler()
	route.GET("/", departmentHandler.GetDepartments)
}
