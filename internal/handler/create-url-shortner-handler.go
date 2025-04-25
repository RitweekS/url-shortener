package handler

import (
	"fmt"
	"os"
	"strings"

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
	baseUrl := os.Getenv("BASEURL")
	if baseUrl == "" {
		fmt.Println("unable to generate URL")
	}

	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("%s/%s", baseUrl, shortUrl),
	})
}

func GetRedirectUrl(ctx *gin.Context) {
	params := ctx.Param("param")
	shortStr := strings.Replace(params, "/", "", 1)

	redirect, err := service.GetRedirectUrl(shortStr)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
	}

	ctx.Redirect(302, redirect)

}
