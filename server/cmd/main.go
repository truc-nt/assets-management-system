package main

import (
	"server/api/routes"
	"server/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	db.ConnectDB()
	routes.LoadRoutes(r)
	r.Run()
}
