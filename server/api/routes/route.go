package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	superGroup := r.Group("/api")
	{
		LoadAssetRoute(superGroup.Group("/assets"))
		LoadDepartmentRoute(superGroup.Group("/departments"))
		LoadAuthRoute(superGroup.Group("/auth"))
	}
}
