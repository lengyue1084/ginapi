package service

import (
	"ginapi/api/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type LoginService struct {
	log *zap.Logger
}

func NewLoginService( log *zap.Logger)*LoginService  {
	return &LoginService{
		log:log,
	}
}

func (l LoginService) Login(ctx *gin.Context)  {
	var json user.UserLoginReq
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	l.log.Info(json.Username + json.Password)
	ctx.JSON(http.StatusOK,json)
}