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

	router := server.NewRouter()

	model.InitDB()

	panic(router.Run(":" + strconv.Itoa(config.ServerPort)))
}
