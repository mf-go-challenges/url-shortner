package main

import (
	"example.com/url-shortner/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handler.RegisterRoutes(server)
	server.Run(":8080")
}
