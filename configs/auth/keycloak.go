package auth

import (
	"context"
	"putra4648/erp/configs/config"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

type KeycloakClaims struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	jwt.RegisteredClaims
}

func GetKeycloakGroups(cfg *config.AppEnv) ([]*gocloak.Group, error) {

	client := gocloak.NewClient(cfg.KeycloakURL)
	ctx := context.Background()

	// 1. Login sebagai Service Account untuk dapat token admin
	token, err := client.LoginClient(ctx, cfg.KeycloakClientID, cfg.KeycloakClientSecret, cfg.KeycloakRealmName)

	// 2. Ambil daftar semua groups
	groups, err := client.GetGroups(ctx, token.AccessToken, cfg.KeycloakRealmName, gocloak.GetGroupsParams{})

	return groups, err
}
