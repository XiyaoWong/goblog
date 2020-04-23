// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// AddPostService .
type AddPostService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Body   string `form:"body" json:"body"`
	IsShow int    `form:"is_show" json:"is_show" binding:"omitempty"`
}

// Add 成功则返回该post模型实例，失败返回错误Response
func (service *AddPostService) Add() (model.Post, *serializer.Response) {
	if service.IsShow != -1 {
		service.IsShow = 1
	}
	post := model.Post{
		Title:  service.Title,
		Body:   service.Body,
		IsShow: service.IsShow,
	}
	// 新增帖子
	if err := model.DB.Create(&post).Error; err != nil {
		return post, &serializer.Response{
			Code:  500,
			Msg:   "新增帖子失败",
			Error: err.Error(),
		}
	}
	return post, nil
}
