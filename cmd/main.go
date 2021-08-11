package main

import (
	"fmt"
	"ginapi/internal/conf"
	"go.uber.org/zap"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "ginApi"
	Version  = "1.0.0"
	flagconf string
	id, _    = os.Hostname()
	logger   = NewZap()
)

func main() {
	conf := conf.NewConf()
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
	return logger
}
