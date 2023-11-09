package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAssetsRoute(route *gin.RouterGroup) {
	assetsHandler := handlers.NewAssetsHandler()
	route.GET("/", assetsHandler.GetAssets)
}
