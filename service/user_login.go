package service

import (
	"goblog/model"
	"goblog/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Code: 401,
			Msg:  "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Code: 401,
			Msg:  "账号或密码错误",
		}
	}
	return user, nil
}
