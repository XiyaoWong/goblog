// Package controller provides ...
package controller

import (
	"goblog/common"
	"goblog/serializer"
	"goblog/service"

	"github.com/gin-gonic/gin"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var service service.AddUserService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Add(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.BuildUserInfoResponse(user))
		}
	} else {
		c.JSON(200, serializer.Response{Code: 400, Msg: "无效数据"})
	}
}

// ShowUserInfo 用户信息
func ShowUserInfo(c *gin.Context) {
	var service service.ShowUserInfoService
	rep := service.Show(c.Param("user_id"))
	c.JSON(200, rep)
}

//UserLogin 登陆
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.Response{Code: 400, Msg: "无效数据"})
		return
	}
	user, errr := service.Login()
	if errr != nil {
		c.JSON(200, errr)
		return
	}
	token, errrr := common.ReleaseToken(user)
	if errrr != nil {
		c.JSON(200, serializer.Response{Code: 500, Msg: "token生成失败"})
		return
	}
	c.JSON(200, serializer.Response{
		Data: gin.H{
			"token": token,
		},
	})
}
