package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .Env")
	}

	return os.Getenv(key)
}
