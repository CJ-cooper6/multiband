package controller

import (
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"multiband/model"
)

func Login(c *gin.Context) {
	var post model.Userlogin
	_ = c.ShouldBindJSON(&post)
	f := model.Flag{2}
	_, err := dao.CheckUsernameAndPassword(post.Username, post.Password)
	if err != nil {
		f.Flag = 0
	} else {
		f.Flag = 1
	}
	c.JSON(200, gin.H{
		"flag": f.Flag,
	})

}
