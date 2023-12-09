package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	superGroup := r.Group("/api")
	{
		LoadAssetRoute(superGroup.Group("/assets"))
		LoadUserRoute(superGroup.Group("/users"))
		LoadAuthRoute(superGroup.Group("/auth"))
	}
}
