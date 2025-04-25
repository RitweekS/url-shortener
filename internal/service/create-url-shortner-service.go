package service

import (
	"math/rand"
	"time"

	"github.com/RitweekS/url-shortener.git/internal/repositories"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	str := make([]byte, 5)
	rand.Seed(time.Now().UnixNano())
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}
	return string(str)
}

func CreateUrlShortener(url string) (string, error) {

	shortenerStr := generateRandomString(5)

	return repositories.CreateUrlShortener(url, shortenerStr)
}

func GetRedirectUrl(shortStr string) (string, error) {
	return repositories.GetRedirectUrl(shortStr)
}
