package main

import (
	"Gin-Learn/config"
	"Gin-Learn/router"
	"fmt"
)

func main() {
	config.Init()
	fmt.Println(config.AppConfig)
	r:=router.InitServer()
	r.Run(":"+config.AppConfig.App.Port)
}