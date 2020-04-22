// Package controller provides ...
package controller

import (
	"goblog/serializer"
	"goblog/service"

	"github.com/gin-gonic/gin"
)

// CheckToken ..
func CheckToken(c *gin.Context) {
	var service service.CheckTokenService
	if err := c.ShouldBind(&service); err == nil {
		rep := service.Check()
		c.JSON(200, rep)
		return
	}
	c.JSON(200, serializer.Response{Msg: "没有token"})
}
