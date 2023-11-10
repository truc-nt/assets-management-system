package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAssetRoute(route *gin.RouterGroup) {
	assetHandler := handlers.NewAssetHandler()
	route.GET("/", assetHandler.GetAssetsByDepartmentId)
	route.GET("/:asset_id", assetHandler.GetAssetById)
	route.POST("/", assetHandler.CreateAsset)
	route.PUT("/:asset_id", assetHandler.UpdateAsset)
	route.DELETE("/:asset_id", assetHandler.DeleteAsset)
}
