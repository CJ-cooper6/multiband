package controller

import (
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"multiband/model"
)

func Regist(c *gin.Context) {
	var user model.User
	var data model.Registuser
	_ = c.ShouldBindJSON(&data)
	flag := model.Flag{2}

	//检查是否有相同的用户名
	uu, _ := dao.CheckUserName(data.Username)

	if uu.Id > 0 {
		//有相同的用户名
		flag.Flag = 0

	} else {
		//没有相同的用户名
		flag.Flag = 1
		user.Username = data.Username
		user.Password = data.Password
		dao.SaveUser(&user)
	}
	c.JSON(200, gin.H{
		"flag": flag.Flag,
	})

}
