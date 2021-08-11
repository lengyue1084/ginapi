// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ginapi/internal/biz"
	"ginapi/internal/conf"
	"ginapi/internal/data"
	"ginapi/internal/service"
	"ginapi/middleware"
	"ginapi/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// initApp init gin application.
func initApp(config *conf.Conf, log *zap.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(middleware.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, router.ProviderSet))
}
