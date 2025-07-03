package main

import (
	"fmt"
	"timeMonitorServer/config"
	"timeMonitorServer/global"
	"timeMonitorServer/router"
)

func main() {
	config.InitEnv()

	err := global.InitRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = global.InitMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Init()
}
