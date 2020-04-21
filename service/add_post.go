// Package service provides ...
package service

import (
	"goblog/model"
	"goblog/serializer"
)

// AddPostService .
type AddPostService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Body  string `form:"body" json:"body"`
}

// Add .
func (service *AddPostService) Add() (model.Post, *serializer.Response) {
	post := model.Post{
		Title: service.Title,
		Body:  service.Body,
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
