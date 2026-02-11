package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	Log = logger.Sugar()

	return logger, nil
}
