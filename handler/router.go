package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/shorten", createShortUrl)
	server.GET("/:code", getOriginalUrl)
	server.POST("/bulk", BulkUploadUrls)
}
