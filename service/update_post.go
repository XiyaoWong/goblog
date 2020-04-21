// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// UpdatePostService ..
type UpdatePostService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Body  string `form:"body" json:"body"`
}

// Update ...
func (service *UpdatePostService) Update(id string) (model.Post, *serializer.Response) {
	var post model.Post
	// 检查是否存在该帖子
	if err := model.DB.Find(&post, id).Error; err != nil {
		return post, &serializer.Response{
			Code:  404,
			Msg:   "帖子不存在",
			Error: err.Error(),
		}
	}
	// 更新帖子
	if err := model.DB.Model(&post).Update("Title", service.Title).Update("Body", service.Body).Error; err != nil {
		return post, &serializer.Response{
			Code:  500,
			Msg:   "新增帖子失败, 情况未知",
			Error: err.Error(),
		}
	}
	return post, nil
}
