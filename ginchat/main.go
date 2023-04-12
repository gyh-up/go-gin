package main

import (
	"go-gin/ginchat/router"
	"go-gin/ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":9090")
}
