package config

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// ServerPort 端口
	ServerPort = 8888
	// ServerMode 模式
	ServerMode = "release"
	// DBName 保存数据库名称
	DBName = "GOBLOG.sqlite3"
	// JWTKey jwt密匙
	JWTKey = "this is awesome!!!"
	// Expriation token有效时间
	Expriation = 3 * 24 * time.Hour
	//AllowRegister 是否开始注册功能
	AllowRegister = false
	// AllowOrigin 允许跨域请求列表
	AllowOrigin []string
)

// Init 初始化配置
func Init() {
	cfg, err := ini.ShadowLoad("config.ini")
	if err != nil {
		fmt.Printf("读取配置文件失败,将按默认配置执行: %v\n", err)
	}
	fmt.Printf("读取配置文件成功\n")
	sec, _ := cfg.GetSection(ini.DEFAULT_SECTION)
	ServerPort = sec.Key("server_port").MustInt(8888)
	ServerMode = sec.Key("server_mode").MustString("release")
	DBName = sec.Key("db_name").MustString("GOBLOG.sqlite3")
	JWTKey = sec.Key("jwt_key").MustString("this is awesome!!!")

	hour := sec.Key("expriation").MustInt(3 * 24)
	Expriation = time.Duration(hour) * time.Hour

	AllowRegister = sec.Key("allow_register").MustBool(false)
	AllowOrigin = sec.Key("origin").ValueWithShadows()
}
