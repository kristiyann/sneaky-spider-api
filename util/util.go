package util

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func LoadEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: '%s'", err)
	}

	result := os.Getenv(key)

	return result
}

func GenerateRandomString(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func MatchRegex(pattern string, text string) (bool, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Println("regex error: " + err.Error())
		return false, err
	}

	return regex.MatchString(text), nil
}
