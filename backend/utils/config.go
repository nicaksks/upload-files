package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}

func Port() string {
	Config()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
		return port
	}
	return ":" + port
}

func Domain() string {
	Config()
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "localhost/"
		return domain
	}
	return domain + "/"
}
