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
	model.InitRedisdb()
	server.InitDealVideo()
	routes.InitRouter()

	//添加缓存
	//a, _, _ := dao.GetVideoallList()
	//for _, val := range a {
	//	str, _ := json.Marshal(val)
	//	model.Redisdb.ZAdd(model.Ctx, "info", &redis.Z{Score: float64(val.ID), Member: str})
	//
	//}

}
