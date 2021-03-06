// Package middleware provides ...
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	myconfig "goblog/config"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.AllowOrigins = myconfig.AllowOrigin
	config.AllowCredentials = true
	return cors.New(config)
}
