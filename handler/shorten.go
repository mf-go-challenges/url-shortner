package handler

import (
	"example.com/url-shortner/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createShortUrl(context *gin.Context) {
	var link models.Link
	err := context.ShouldBindJSON(&link)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	link.UserID = userId

	ShortUrl, err := link.ShortenUrl()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, ShortUrl)
}

func getOriginalUrl(context *gin.Context) {
	code := context.Param("code")
	link := &models.Link{Code: code}
	err := link.GetOriginalUrl()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Redirect(302, link.Url)
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

	var link models.Link
	userId := context.GetInt64("userId")
	link.UserID = userId

	result, err := link.BulkUploadUrls(openedFile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, result)
}
