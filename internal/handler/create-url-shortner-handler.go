package handler

import (
	"fmt"

	"github.com/RitweekS/url-shortener.git/internal/dto/requests"
	"github.com/RitweekS/url-shortener.git/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateUrlShortener(ctx *gin.Context) {
	var requestBody requests.CreateUrlShortener

	jsonBindErr := ctx.BindJSON(&requestBody)

	if jsonBindErr != nil {
		ctx.JSON(200, gin.H{
			"message": jsonBindErr.Error(),
		})
		return
	}
	shortUrl, err := service.CreateUrlShortener(requestBody.Url)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	ctx.Redirect(302, fmt.Sprintf("http://localhost:3000/%s", shortUrl))
}
