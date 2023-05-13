package routes

import (
	"github.com/gin-gonic/gin"
	"multiband/controller"
	"multiband/middleware"
	"multiband/utils"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/hello", controller.Hello)
	_ = r.Run(utils.HttpPort)
}
