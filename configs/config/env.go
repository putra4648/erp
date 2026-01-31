package config

import (
	"os"
	"putra4648/erp/configs/logger"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	KeycloakURL          string
	KeycloakClientID     string
	KeycloakClientSecret string
	KeycloakRealmName    string
	DBDSN                string
	Port                 string
}

func LoadConfig() *AppEnv {
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Log.Fatalf("Error loading .env file: %v", err)
	}

	return &AppEnv{
		KeycloakURL:          getEnv("KEYCLOAK_URL", ""),
		KeycloakClientID:     getEnv("KEYCLOAK_CLIENT_ID", ""),
		KeycloakClientSecret: getEnv("KEYCLOAK_CLIENT_SECRET", ""),
		KeycloakRealmName:    getEnv("KEYCLOAK_REALM_NAME", ""),
		DBDSN:                getEnv("DB_DSN", ""),
		Port:                 getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
