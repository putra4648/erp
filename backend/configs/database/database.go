package database

import (
	"putra4648/erp/configs/config"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.AppEnv, zapLogger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DBDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		zapLogger.Sugar().Fatalf("Cannot connect to DB: %v", err)
	}

	return db, err
}
