package middlewares

import (
	"example.com/url-shortner/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header required"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized", "error": err.Error()})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
