// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// AddUserService 管理用户注册服务
type AddUserService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Avatar          string `form:"avatar" json:"avatar" binding:"required"`
}

// Valid 验证数据
func (service *AddUserService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 401,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 401,
			Msg:  "用户名已存在",
		}
	}
	return nil
}

// Add 用户注册
func (service *AddUserService) Add() (model.User, *serializer.Response) {
	user := model.User{
		Nickname: service.Nickname,
		Username: service.Username,
		Avatar:   service.Avatar,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Code: 500,
			Msg:  "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Code: 500,
			Msg:  "注册失败",
		}
	}

	return user, nil
}
