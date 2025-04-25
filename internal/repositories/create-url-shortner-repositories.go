package repositories

import (
	"fmt"

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

func GetRedirectUrl(shortStr string) (string, error) {
	var result struct {
		Url string
	}
	err := database.DB.Model(&models.Shortener{}).
		Select("url").
		Where("short_str = ?", shortStr).
		Scan(&result).Error

	if err != nil {
		fmt.Println("Error:", err)
	}

	return result.Url, nil
}
