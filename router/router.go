package router

import (
	"ginapi/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)
var ProviderSet = wire.NewSet(NewApp,NewRouter)

func NewRouter(
	app *App,
	UserServer *service.UserServer,
) *gin.Engine {
	router := app.app.Group("/")
	{
		router.GET("/login", UserServer.Login)
	}
	return app.app
}