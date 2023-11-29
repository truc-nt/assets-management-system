package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadDepartmentRoute(route *gin.RouterGroup) {
	employeeHandler := handlers.NewEmployeeHandler()
	route.GET("/", employeeHandler.GetEmployees)
}
