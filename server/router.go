// Package server provides ...
package server

import (
	"goblog/config"
	"goblog/controller"
	"goblog/middleware"
	"goblog/serializer"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.SERVER_MODE)
	r := gin.Default()

	// 跨域
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) { c.JSON(200, serializer.Response{Msg: "Pong!"}) })

	if config.ALLOW_REGISTER {
		r.POST("/users", controller.AddUser)
	}

	r.POST("/login", controller.UserLogin)

	r.GET("users/:user_id", controller.ShowUserInfo)

	r.GET("/posts", controller.ShowPostList)
	r.GET("/posts/:post_id", controller.ShowPostDetail)

	// 需要鉴权的
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/posts", controller.AddPost)
	auth.PUT("/posts/:post_id", controller.UpdatePost)
	auth.DELETE("/posts/:post_id", controller.DeletePost)

	return r
}
