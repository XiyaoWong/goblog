package config

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	SERVER_PORT    = 8888
	SERVER_MODE    = "release"
	DB_NAME        = "GOBLOG.sqlite3"
	JWT_KEY        = "this is awesome!!!"
	EXPRIATION     = 3 * 24 * time.Hour
	ALLOW_REGISTER = false
	ALLOW_ORIGIN   []string
)

func Init() {
	cfg, err := ini.ShadowLoad("config.ini")
	if err != nil {
		fmt.Printf("读取配置文件失败,将按默认配置执行: %v\n", err)
	}
	fmt.Printf("读取配置文件成功\n")
	sec, _ := cfg.GetSection(ini.DEFAULT_SECTION)
	SERVER_PORT = sec.Key("server_port").MustInt(8888)
	SERVER_MODE = sec.Key("server_mode").MustString("release")
	DB_NAME = sec.Key("db_name").MustString("GOBLOG.sqlite3")
	JWT_KEY = sec.Key("jwt_key").MustString("this is awesome!!!")

	hour := sec.Key("expriation").MustInt(3 * 24)
	EXPRIATION = time.Duration(hour) * time.Hour

	ALLOW_REGISTER = sec.Key("allow_register").MustBool(false)
	ALLOW_ORIGIN = sec.Key("origin").ValueWithShadows()
}
