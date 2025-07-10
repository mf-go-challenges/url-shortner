package handler

import (
	"example.com/url-shortner/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
