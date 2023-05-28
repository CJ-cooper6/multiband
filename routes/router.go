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
	r.POST("/regist", controller.Regist)
	r.POST("/login", controller.Login)
	auth := r.Group("user")
	{
		auth.GET("/hello", controller.Hello)
		auth.GET("/download", controller.Download)
		auth.POST("/del", controller.Del)
		auth.POST("/play", controller.Play)
	}

	_ = r.Run(utils.HttpPort)
}
