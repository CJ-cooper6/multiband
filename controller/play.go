package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"multiband/server"
	"multiband/utils"
	"os"
	"path/filepath"
	"strconv"
)

func Play(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	meta, err := dao.GetVideoMeta(id)
	if err != nil {
		fmt.Println(err)
	}

	// 转化格式
	outputFile := utils.Trans + "/" + filepath.Base(meta.Location[:len(meta.Location)-4]) + ".mp4"
	if ok, _ := exists(outputFile); ok {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  200,
			"path": utils.Trans + "/" + filepath.Base(meta.Location[:len(meta.Location)-4]) + ".mp4",
		})
		return
	} else {
		err = server.TransVideo(meta.Location, outputFile)
		if err != nil {
			fmt.Println(err)
		}
	}

	//返回视频路径
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  200,
		"path": utils.Trans + "/" + filepath.Base(meta.Location[:len(meta.Location)-4]) + ".mp4",
	})

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
