package main

import (
	"example.com/url-shortner/db"
	"example.com/url-shortner/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	handler.RegisterRoutes(server)
	server.Run(":8080")
}
