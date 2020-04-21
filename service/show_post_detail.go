// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// ShowPostDetailService .
type ShowPostDetailService struct{}

// Show .
func (service *ShowPostDetailService) Show(id string) serializer.Response {
	var post model.Post
	err := model.DB.First(&post, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "帖子不存在",
			Error: err.Error(),
		}
	}
	return serializer.BuildPostDetailResponse(post)

}
