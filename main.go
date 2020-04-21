// Package main provides ...
package main

import (
	"goblog/config"
	"goblog/model"
	"goblog/server"
	"strconv"
)

func main() {
	config.Init()
	model.InitDB()

	router := server.NewRouter()

	panic(router.Run(":" + strconv.Itoa(config.ServerPort)))
}
