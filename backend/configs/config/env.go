package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type AppEnv struct {
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0Domain       string
	Auth0Audience     string
	DBDSN             string
	Port              string
}

func LoadConfig(zapLogger *zap.Logger) *AppEnv {
	// Try to load .env from several possible locations relative to the binary/execution context.
	// We don't fatal here because environment variables might already be set in the system/docker.
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../../.env")

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
