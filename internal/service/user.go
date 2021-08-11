package service

import (
	"fmt"
	"ginapi/internal/biz"
	"github.com/gin-gonic/gin"
)

type UserServer struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserServer {
	return &UserServer{
		uc: uc,
	}
}

func (u *UserServer) Login(ctx *gin.Context) {
	fmt.Println(111111111111)
	u.uc.UserLogin(ctx)
}
