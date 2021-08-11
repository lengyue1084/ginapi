package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMiddleware)

type middleware interface {
	Cors() gin.HandlerFunc
	Logger() gin.HandlerFunc
	Recovery() gin.HandlerFunc
}

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}
