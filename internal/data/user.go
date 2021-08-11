package data

import (
	"fmt"
	"ginapi/api/user"
	"ginapi/internal/biz"
	"ginapi/internal/data/model"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

type userRepo struct {
	UserData *UserData
}

func NewUserRepo(d *UserData) biz.UserRepo {
	return &userRepo{
		UserData: d,
	}
}

func (u *userRepo) Login(ctx *gin.Context) (*user.LoginReplay, error) {
	userMode := &model.User{
		Name: "WQ",
		Age:  18,
	}
	res := u.UserData.db.Create(&userMode)
	fmt.Println(res)
	var user1 []model.User
	result := u.UserData.db.Table("users").Find(&user1)
	fmt.Println(reflect.TypeOf(&user1))
	fmt.Println(result)
	fmt.Println(userMode)
	var list = user.LoginReplay{
		Code:    0,
		Message: "success",
	}
	//err := rdb.Set(ctx, "key", "value", 0).Err()
	u.UserData.redis.Set(ctx, "ginapi:user01", "WQredis", 10*time.Second)
	key := u.UserData.redis.Get(ctx, "ginapi:user01")
	fmt.Println(key)
	u.UserData.log.Info("111111111111")

	return &list, nil
}
