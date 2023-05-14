package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"strconv"
)

func Hello(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}
	meta, total, err := dao.GetVideoMeta(page, limit)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "",
		"count": total,
		"data":  meta,
	})

}
