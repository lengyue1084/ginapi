package main

import (
	"flag"
	"fmt"
	"ginapi/internal/conf"
	"go.uber.org/zap"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "ginApi"
	Version  = "1.0.0"
	flagConf string
	id, _    = os.Hostname()
	logger   *zap.Logger
)

func init() {
	flag.StringVar(&flagConf, "conf", "./configs", "config path, eg: -conf config.yaml")
}
func main() {
	flag.Parse()
	config := conf.NewConf(flagConf)
	logger = conf.NewZap()
	app, cleanup, err := initApp(config, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// 启动服务
	if err := app.Run(
		fmt.Sprintf("%s:%s", config.Conf.GetString("server.http.addr"),
			config.Conf.GetString("server.http.port")),
	); err != nil {
		panic(err)
	}
}
