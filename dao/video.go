package dao

import (
	"fmt"
	"multiband/model"
)

//存储视频信息
func SaveVideoMeta(meta *model.VideoMeta) {
	err := model.Db.Create(meta).Error
	if err != nil {
		fmt.Println(err)
	}
}

//
func GetVideoMeta(page, limit int) ([]model.VideoMeta, int64, error) {
	var meta []model.VideoMeta
	var total int64
	err := model.Db.Limit(limit).Offset((page - 1) * limit).Find(&meta).Count(&total).Error
	if err != nil {
		return meta, 0, err
	}
	return meta, total, nil
}
