package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type data struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	ImgPath  string `json:"imgPath"`
}

func Hello(c *gin.Context) {
	da := []data{
		{
			Id:       1,
			Username: "zhangsan",
			ImgPath:  "../imgs/1.png",
		},
		{
			Id:       2,
			Username: "lisi",
			ImgPath:  "../imgs/2.png",
		},
		{
			Id:       3,
			Username: "lisi",
		},
		{
			Id:       4,
			Username: "lisi",
		},
	}
	a, _ := json.Marshal(da)
	fmt.Println(string(a))
	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "",
		"count": 5,
		"data":  da,
	})

}
