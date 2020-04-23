// Package service provides ...
package service

import (
	"fmt"
	"goblog/common"
	"goblog/model"
	"goblog/serializer"
)

// ShowPostListService ..
type ShowPostListService struct{}

// Show ...
func (service *ShowPostListService) Show(tokenString string) *serializer.Response {
	posts := []model.Post{}

	fmt.Printf("tokenString: %v", tokenString)
	token, _, err := common.ParseToken(tokenString)
	// 验证不通过 仅显示公开内容
	if err != nil || !token.Valid {
		fmt.Println("验证不通过")
		if err := model.DB.Where("is_show = ?", 1).Find(&posts).Error; err != nil {
			return &serializer.Response{
				Code:  500,
				Msg:   "服务器数据库查询错误",
				Error: err.Error(),
			}
		}
	} else {
		// 验证通过 全部显示
		fmt.Println("验证通过")
		if err := model.DB.Find(&posts).Error; err != nil {
			return &serializer.Response{
				Code:  500,
				Msg:   "服务器数据库查询错误",
				Error: err.Error(),
			}
		}
	}

	return &serializer.Response{
		Data: serializer.BuildPostList(posts),
	}
}
