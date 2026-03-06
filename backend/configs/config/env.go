package config

import (
	"os"
	"putra4648/erp/configs/logger"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0Domain       string
	Auth0Audience     string
	DBDSN             string
	Port              string
}

func LoadConfig() *AppEnv {
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Log.Fatalf("Error loading .env file: %v", err)
	}

	return &AppEnv{
		Auth0ClientID:     getEnv("AUTH0_CLIENT_ID", ""),
		Auth0ClientSecret: getEnv("AUTH0_CLIENT_SECRET", ""),
		Auth0Domain:       getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience:     getEnv("AUTH0_AUDIENCE", ""),
		DBDSN:             getEnv("DB_DSN", ""),
		Port:              getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
