package biz

import (
	"ginapi/api/user"
	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	Login(ctx *gin.Context) (*user.LoginReplay, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) UserLogin(ctx *gin.Context) (*user.LoginReplay, error){
	return u.repo.Login(ctx)
	//return nil,nil

}
