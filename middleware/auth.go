// Package middleware provides ...
package middleware

import (
	"fmt"
	"goblog/common"
	"goblog/model"
	"goblog/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(200, serializer.Response{Code: 401, Msg: "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(200, serializer.Response{Code: 401, Msg: "权限不足"})
			fmt.Println(err)
			c.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userID := claims.UserID
		var user model.User
		model.DB.First(&user, userID)

		// 用户
		if user.ID == 0 {
			c.JSON(200, serializer.Response{Code: 401, Msg: "权限不足"})
			c.Abort()
			return
		}

		// 用户存在 将user 的信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
