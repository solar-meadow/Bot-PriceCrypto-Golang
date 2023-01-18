package token

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func BotTokenFromEnv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	token := os.Getenv(key)
	if len(token) == 0 {
		return "", errors.New("bot token not found from .env")
	}
	return token, nil
}
