package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) Logger() gin.HandlerFunc {
	return gin.Logger()
}

func (m *Middleware) Recovery() gin.HandlerFunc {
	return gin.Recovery()
}
