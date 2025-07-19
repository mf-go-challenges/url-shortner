package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"net/http"
)

func RateLimitermiddleware(limiter *redis_rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		clientIP := c.ClientIP()
		res, err := limiter.Allow(ctx, clientIP, redis_rate.PerMinute(10))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if res.Remaining == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			return
		}

		c.Next()
	}
}
