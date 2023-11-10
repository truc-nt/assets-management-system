package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAuthRoute(route *gin.RouterGroup) {
	authHandler := handlers.NewAuthHandler()
	route.POST("/login", authHandler.Login)
	route.POST("/logout", authHandler.Logout)
}
