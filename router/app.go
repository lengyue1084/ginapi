package router

import (
	"ginapi/middleware"
	"github.com/gin-gonic/gin"
)
type App struct {
	app *gin.Engine
}

func NewApp(
	middleware *middleware.Middleware,
) *App {
	r := gin.New()
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	r.Use(middleware.Cors(),middleware.Logger(),middleware.Recovery())
	return &App{
		app: r,
	}
}