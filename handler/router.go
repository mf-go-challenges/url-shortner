package handler

import (
	"example.com/url-shortner/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RegisterRoutes(server *gin.Engine, limiter *redis_rate.Limiter) {
	server.POST("/shorten", middlewares.Authenticate, middlewares.RateLimitermiddleware(limiter), createShortUrl)
	server.POST("/bulk", middlewares.Authenticate, middlewares.RateLimitermiddleware(limiter), BulkUploadUrls)

	server.GET("/:code", getOriginalUrl)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
