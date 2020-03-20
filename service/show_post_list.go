// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

type ShowPostListService struct{}

func (service *ShowPostListService) Show() serializer.Response {
	posts := []model.Post{}
	if err := model.DB.Find(&posts).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器数据库查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildPostList(posts),
	}
}
