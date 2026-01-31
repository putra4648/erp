package auth

import (
	"putra4648/erp/configs/config"

	"github.com/coreos/go-oidc/v3/oidc"
)

func NewOIDCVerifier(cfg *config.AppEnv, provider *oidc.Provider) *oidc.IDTokenVerifier {
	return provider.Verifier(&oidc.Config{ClientID: cfg.KeycloakClientID})
}
