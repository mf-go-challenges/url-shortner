package main

import (
	"example.com/url-shortner/db"
	"example.com/url-shortner/handler"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	db.InitDB()
	server := gin.Default()
	handler.RegisterRoutes(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
