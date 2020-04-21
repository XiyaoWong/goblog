// Package controller provides ...
package controller

import (
	"goblog/service"

	"goblog/serializer"

	"github.com/gin-gonic/gin"
)

// AddPost 新增帖子
func AddPost(c *gin.Context) {
	var service service.AddPostService
	if err := c.ShouldBind(&service); err == nil {
		if post, err := service.Add(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.BuildPostDetailResponse(post))
		}
	} else {
		c.JSON(200, serializer.Response{Code: 400, Msg: "无效数据"})
	}
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	var service service.DeletePostService
	rep := service.Delete(c.Param("post_id"))
	c.JSON(200, rep)

}

// ShowPostList 显示帖子列表
func ShowPostList(c *gin.Context) {
	var service service.ShowPostListService
	rep := service.Show()
	c.JSON(200, rep)
}

// ShowPostDetail 显示帖子详情
func ShowPostDetail(c *gin.Context) {
	var service service.ShowPostDetailService
	rep := service.Show(c.Param("post_id"))
	c.JSON(200, rep)
}

// UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	var service service.UpdatePostService
	if err := c.ShouldBind(&service); err == nil {
		if post, err := service.Update(c.Param("post_id")); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.BuildPostDetailResponse(post))
		}
	} else {
		c.JSON(200, serializer.Response{Code: 400, Msg: "无效数据"})
	}
}
