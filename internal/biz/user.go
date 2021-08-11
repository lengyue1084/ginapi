package biz

import (
	"ginapi/api/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRepo interface {
	Login(ctx *gin.Context) (*user.LoginReplay, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *zap.Logger
}

func NewUserUseCase(repo UserRepo, log *zap.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log}
}

func (u *UserUseCase) UserLogin(ctx *gin.Context) (*user.LoginReplay, error) {
	u.log.Info("biz 示例")
	return u.repo.Login(ctx)
}
