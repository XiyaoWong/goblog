// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// ShowUserInfoService ...
type ShowUserInfoService struct{}

// Show ...
func (service *ShowUserInfoService) Show(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "用户不存在",
			Error: err.Error(),
		}
	}
	return serializer.BuildUserInfoResponse(user)

}
