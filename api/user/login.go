package user

type UserLoginReq struct {
	Username     string `form:"user" json:"username" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}