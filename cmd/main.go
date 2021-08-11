package main

import (
	"fmt"
	"ginapi/internal/conf"
	"go.uber.org/zap"
	"os"
	"time"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "ginApi"
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
	Logger  = NewZap()
)

func main() {
	conf := conf.NewConf()
	logger := NewZap()
	logger.Info("safasfsfsafsafasfas")
	app, cleanup, err := initApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// 启动服务
	if err := app.Run(
		fmt.Sprintf("%s:%s", conf.Conf.GetString("server.http.addr"),
			conf.Conf.GetString("server.http.port")),
	); err != nil {
		panic(err)
	}
}

func NewZap() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	return logger
}
