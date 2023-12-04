package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadUserRoute(route *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()
	route.GET("/", userHandler.GetUsers)
	route.GET("/:user_id", userHandler.GetUserById)
	route.PUT("/:user_id", userHandler.UpdateUser)
}
