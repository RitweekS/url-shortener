package internal

import (
	"github.com/RitweekS/url-shortener.git/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/short", handler.CreateUrlShortener)
	router.GET("/:param", handler.GetRedirectUrl)
}
