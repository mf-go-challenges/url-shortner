package handler

import (
	"example.com/url-shortner/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createShortUrl(context *gin.Context) {
	var url models.ShortUrl
	err := context.ShouldBindJSON(&url)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "error": err.Error()})
		return
	}

	ShortUrl, err := url.ShortenUrl()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, ShortUrl)
}

func getOriginalUrl(context *gin.Context) {
	code := context.Param("code")
	url, exists := models.UrlStore[code]
	if !exists {
		context.JSON(http.StatusNotFound, gin.H{"error": "code not found"})
		return
	}

	context.Redirect(302, url)
}

func BulkUploadUrls(context *gin.Context) {
	fileHeader, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	openedFile, err := fileHeader.Open()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "cannot open file"})
		return
	}

	defer openedFile.Close()

	result, err := models.BulkUploadUrls(openedFile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, result)
}
