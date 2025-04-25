package repositories

import (
	"github.com/RitweekS/url-shortener.git/internal/database"
	"github.com/RitweekS/url-shortener.git/internal/models"
)

func CreateUrlShortener(redirectUrl string, shortenerStr string) (string, error) {
	payload := models.Shortener{
		Url:      redirectUrl,
		ShortStr: shortenerStr,
	}
	created := database.DB.Create(&payload)

	if created.Error != nil {
		return "", created.Error
	}
	return payload.ShortStr, nil
}
