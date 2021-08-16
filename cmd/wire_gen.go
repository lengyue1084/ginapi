// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"ginapi/internal/biz"
	"ginapi/internal/conf"
	"ginapi/internal/data"
	"ginapi/internal/service"
	"ginapi/middleware"
	"ginapi/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// initApp init gin application.
func initApp(config *conf.Conf, log *zap.Logger) (*gin.Engine, func(), error) {
	middlewareMiddleware := middleware.NewMiddleware()
	app := router.NewApp(middlewareMiddleware, config)
	db, cleanup, err := data.NewGormClient(config)
	if err != nil {
		return nil, nil, err
	}
	client, cleanup2, err := data.NewRedisClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userData, cleanup3, err := data.NewUserData(db, client, log)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(userData)
	userUseCase := biz.NewUserUseCase(userRepo, log)
	userService := service.NewUserService(userUseCase, log)
	homeService := service.NewHome(log)
	engine := router.NewRouter(app, userService, homeService)
	return engine, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
