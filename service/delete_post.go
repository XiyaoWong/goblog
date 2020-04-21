// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// DeletePostService .
type DeletePostService struct{}

// Delete .
func (service *DeletePostService) Delete(id string) serializer.Response {
	var post model.Post
	err := model.DB.First(&post, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "帖子不存在",
			Error: err.Error(),
		}
	}
	err = model.DB.Delete(&post).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "服务器错误，处理结果未知",
		}
	}
	return serializer.Response{
		Msg: "删除成功",
	}
}
