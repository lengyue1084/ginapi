package service

import (
	"ginapi/internal/biz"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	uc  *biz.UserUseCase
	log *zap.Logger
}

func NewUserService(uc *biz.UserUseCase, log *zap.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log,
	}
}

func (u *UserService) Login(ctx *gin.Context) {
	u.log.Info("service示例")
	u.uc.UserLogin(ctx)
}
