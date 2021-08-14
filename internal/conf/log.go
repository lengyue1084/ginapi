package conf

import "go.uber.org/zap"

func NewZap() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger
}
