package auth

import (
	"context"
	"putra4648/erp/configs/config"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	Audience string
}

func NewAuthenticator(cfg *config.AppEnv) (*Authenticator, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://"+cfg.Auth0Domain+"/")
	if err != nil {
		return nil, err
	}

	return &Authenticator{
		Provider: provider,
		Config: oauth2.Config{
			ClientID:     cfg.Auth0ClientID,
			ClientSecret: cfg.Auth0ClientSecret,
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
			Endpoint:     provider.Endpoint(),
		},
		Audience: cfg.Auth0Audience,
	}, nil
}

func (a *Authenticator) VerifyToken(ctx context.Context, tokenString string) (*oidc.IDToken, error) {
	oidcConfig := &oidc.Config{
		ClientID: a.Audience,
	}

	return a.Verifier(oidcConfig).Verify(ctx, tokenString)
}
