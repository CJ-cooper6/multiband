package main

import (
	"multiband/model"
	"multiband/routes"
	"multiband/server"
	"multiband/utils"
)

func main() {
	utils.Setting() //初始化
	model.IntDb()
	server.InitDealVideo()
	routes.InitRouter()
}
