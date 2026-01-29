package main

import (
	"putra4648/erp/configs/config"
	supplierService "putra4648/erp/internal/modules/inventory/supplier/service"
	warehouseService "putra4648/erp/internal/modules/inventory/warehouse/service"

	"github.com/casbin/casbin/v3"
	"github.com/coreos/go-oidc/v3/oidc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppDependencies struct {
	Config                  *config.AppEnv
	DB                      *gorm.DB
	Enforcer                *casbin.Enforcer
	Verifier                *oidc.IDTokenVerifier
	ZapLogger               *zap.Logger
	WarehouseCommandService warehouseService.WarehouseCommandService
	WarehouseQueryService   warehouseService.WarehouseQueryService
	SupplierCommandService  supplierService.SupplierCommandService
	SupplierQueryService    supplierService.SupplierQueryService
}
