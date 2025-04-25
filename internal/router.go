package internal

import (
	"github.com/RitweekS/url-shortener.git/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/v1")
	{
		api.POST("/short", handler.CreateUrlShortener)
	}
}
