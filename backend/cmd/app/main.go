package main

import (
	"fmt"
	"log"
	app "putra4648/erp/cmd"
	"putra4648/erp/configs/auth"
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/database"
	"putra4648/erp/configs/logger"
	"putra4648/erp/internal/category"
	"putra4648/erp/internal/product"
	"putra4648/erp/internal/stock_adjustment"
	"putra4648/erp/internal/stock_level"
	"putra4648/erp/internal/stock_movement"
	"putra4648/erp/internal/supplier"
	"putra4648/erp/internal/uom"
	"putra4648/erp/internal/warehouse"

	"go.uber.org/dig"
)

type Dependency struct {
	Constructror interface{}
	Token        string
}

func main() {
	container := dig.New()

	// Register constructors without names if Token is empty
	if err := container.Provide(logger.InitLogger); err != nil {
		fmt.Printf("Failed to provide logger: %v\n", err)
	}
	if err := container.Provide(config.LoadConfig); err != nil {
		fmt.Printf("Failed to provide config: %v\n", err)
	}
	if err := container.Provide(database.InitDatabase); err != nil {
		fmt.Printf("Failed to provide database: %v\n", err)
	}
	if err := container.Provide(auth.NewAuthenticator); err != nil {
		fmt.Printf("Failed to provide authenticator: %v\n", err)
	}

	// For modules, if they fail to register, we use standard log because
	// logger.Log might not be initialized by dig yet (constructors run only on Invoke)
	// although our init() should have filled it, standard log is safer here.
	if err := warehouse.Register(container); err != nil {
		log.Fatalf("Failed to register warehouse module: %v", err)
	}

	if err := supplier.Register(container); err != nil {
		log.Fatalf("Failed to register supplier module: %v", err)
	}

	if err := product.Register(container); err != nil {
		log.Fatalf("Failed to register product module: %v", err)
	}

	if err := category.Register(container); err != nil {
		log.Fatalf("Failed to register category module: %v", err)
	}

	if err := uom.Register(container); err != nil {
		log.Fatalf("Failed to register uom module: %v", err)
	}

	if err := stock_adjustment.Register(container); err != nil {
		log.Fatalf("Failed to register stock_adjustment module: %v", err)
	}

	if err := stock_movement.Register(container); err != nil {
		log.Fatalf("Failed to register stock_movement module: %v", err)
	}

	if err := stock_level.Register(container); err != nil {
		log.Fatalf("Failed to register stock_level module: %v", err)
	}

	if err := container.Invoke(app.Server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
