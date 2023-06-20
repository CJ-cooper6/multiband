package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"multiband/model"
	"strconv"
)

func Hello(c *gin.Context) {
	var meta []model.VideoMeta
	var total int64
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}
	////先查缓存
	//val, _ := model.Redisdb.ZRange(model.Ctx, "info", int64((page-1)*limit), int64((page-1)*limit+limit)).Result()
	//if len(val) > 0 {
	//	for _, v := range val {
	//		a := model.VideoMeta{}
	//		json.Unmarshal([]byte(v), &a)
	//		meta = append(meta, a)
	//	}
	//
	//	c.JSON(200, gin.H{
	//		"code":  0,
	//		"msg":   200,
	//		"count": total,
	//		"data":  meta,
	//	})
	//	return
	//}

	meta, total, err := dao.GetVideoList(page, limit)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code":  0,
		"msg":   200,
		"count": total,
		"data":  meta,
	})

}
