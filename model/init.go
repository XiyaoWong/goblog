// Package model provides ...
package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"goblog/config"
)

// DB ...
var DB *gorm.DB

// InitDB ...
func InitDB() {
	db, err := gorm.Open("sqlite3", config.DBName)
	db.LogMode(true)
	if err != nil {
		panic("数据库连接错误")
	}
	db.LogMode(gin.Mode() != gin.ReleaseMode)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
