package service

import (
	"ginapi/internal/biz"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserServer struct {
	uc  *biz.UserUseCase
	log *zap.Logger
}

func NewUserService(uc *biz.UserUseCase, log *zap.Logger) *UserServer {
	return &UserServer{
		uc:  uc,
		log: log,
	}
}

func (u *UserServer) Login(ctx *gin.Context) {
	u.log.Info("service示例")
	u.uc.UserLogin(ctx)
}
