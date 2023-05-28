package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"strconv"
)

func Del(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	err := dao.DelVideoMeta(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  500,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  200,
		})
	}

}
