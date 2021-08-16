package router

import (
	"ginapi/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewApp, NewRouter)

func NewRouter(
	app *App,
	UserServer *service.UserService,
	HomeServer *service.HomeService,
) *gin.Engine {
	router := app.app.Group("/")
	{
		router.GET("/", HomeServer.Index)
		router.GET("/login", UserServer.Login)
	}
	return app.app
}
