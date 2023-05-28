package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"multiband/dao"
	"path/filepath"
	"strconv"
)

func Download(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	meta, err := dao.GetVideoMeta(id)
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(meta.Location))
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(meta.Location)
}
