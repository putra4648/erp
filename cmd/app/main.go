package main

import (
	"fmt"
	app "putra4648/erp/cmd"
	"putra4648/erp/configs/auth"
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/database"
	"putra4648/erp/configs/logger"
	"putra4648/erp/internal/modules/inventory"

	"go.uber.org/dig"
)

type Dependency struct {
	Constructror interface{}
	Token        string
}

func main() {
	container := dig.New()

	// Register constructors
	providers := []Dependency{
		{
			Constructror: logger.InitLogger,
			Token:        "",
		},
		{
			Constructror: config.LoadConfig,
			Token:        "",
		},
		{
			Constructror: database.InitDatabase,
			Token:        "",
		},
		{
			Constructror: auth.SetupCasbin,
			Token:        "",
		},
		{
			Constructror: auth.NewOIDCProvider,
			Token:        "",
		},
		{
			Constructror: auth.NewOIDCVerifier,
			Token:        "",
		},
	}

	for _, p := range providers {
		if err := container.Provide(p.Constructror, dig.Name(p.Token)); err != nil {
			fmt.Printf("Failed to initialize %v", err)
		}
	}

	if err := inventory.Register(container); err != nil {
		logger.Log.Fatalf("Failed to register inventory module: %v", err)
	}

	if err := container.Invoke(app.Server); err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
	}

}
