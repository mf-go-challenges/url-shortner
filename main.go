package main

import (
	"os"

	"example.com/url-shortner/db"
	"example.com/url-shortner/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(cors.Default())
	handler.RegisterRoutes(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
