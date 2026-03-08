package logger

import (
	"log"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	// Initialize with a development logger by default to avoid nil pointer issues
	// if Log is called before InitLogger.
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("Failed to initialize default logger: %v", err)
		return
	}
	Log = logger.Sugar()
}

func InitLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	Log = logger.Sugar()

	return logger, nil
}
