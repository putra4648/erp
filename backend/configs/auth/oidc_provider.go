package auth

import (
	"context"
	"putra4648/erp/configs/config"

	"github.com/coreos/go-oidc/v3/oidc"
)

func NewOIDCProvider(cfg *config.AppEnv) (*oidc.Provider, error) {
	return oidc.NewProvider(context.Background(), cfg.KeycloakURL+"/realms/"+cfg.KeycloakRealmName)
}
