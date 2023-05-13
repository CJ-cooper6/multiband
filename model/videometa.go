package model

import "gorm.io/gorm"

type VideoMeta struct {
	gorm.Model
	FileSha1         string `gorm:"type:varchar(20);not null;unique;" json:"file_sha_1"` //文件哈希值
	FileName         string `gorm:"type:varchar(20);not null;" json:"file_name"`         // 文件名称
	FileSize         int64  `gorm:"type:varchar(20);not null;" json:"file_size"`         //文件大小
	Location         string `gorm:"type:varchar(20);not null;" json:"location"`          //视频存放位置
	Picture_Location string `gorm:"type:varchar(20);not null;" json:"picture_location"`  //视频图片存放位置
}
