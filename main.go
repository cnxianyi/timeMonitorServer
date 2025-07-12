package main

import (
	"fmt"
	"timeMonitorServer/config"
	"timeMonitorServer/global"
	"timeMonitorServer/router"
	"timeMonitorServer/utils"
)

func main() {
	config.InitEnv()

	err := global.InitMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.InitCronJobs()

	router.Init()
}
