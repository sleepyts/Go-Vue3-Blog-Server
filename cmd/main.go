package main

import (
	"Go-Vue3-Blog-Server/config"
	"Go-Vue3-Blog-Server/router"
)

func main() {
	config.Init()

	r := router.InitServer()

	r.Run(":" + config.AppConfig.App.Port)
}
