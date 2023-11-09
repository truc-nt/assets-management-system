package main

import (
	"log"
	"server/api/routes"
	"server/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	db.ConnectDB()
	routes.LoadRoutes(r)
	r.Run()
}
