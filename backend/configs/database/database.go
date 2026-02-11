package database

import (
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.AppEnv) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DBDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Log.Fatalf("Cannot connect to DB: %v", err)
	}

	return db, err
}
