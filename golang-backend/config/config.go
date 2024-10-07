package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	DB_URL            string
	PORT              string
	HasuraEndpoint    string // Exported variable
	HasuraAdminSecret string // Exported variable
	JWT_SECRET_KEY    string
	FROM              string
	EMAIL_HOST        string
	EMAIL_PASSWORD    string
	EMAIL_PORT        int
)

func LoadConfig() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	DB_URL = os.Getenv("DB_URL")
	PORT = os.Getenv("PORT")
	HasuraEndpoint = os.Getenv("HASURA_ENDPOINT")        // Adjust the case to match the environment variable
	HasuraAdminSecret = os.Getenv("HASURA_ADMIN_SECRET") // Adjust the case to match the environment variable
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")         // Adjust the case to match the environment variable
	FROM = os.Getenv("FROM")
	EMAIL_HOST = os.Getenv("EMAIL_HOST")
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	EMAIL_PORT, err = strconv.Atoi(os.Getenv("EMAIL_PORT"))
}
