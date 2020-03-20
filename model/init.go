// Package model provides ...
package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"goblog/config"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", config.DB_NAME)
	db.LogMode(true)
	if err != nil {
		panic("数据库连接错误")
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
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
