package model

import (
	"gorm.io/gorm"
	"time"
)

type VideoMeta struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	FileName         string    `gorm:"type:varchar(100);not null;" json:"file_name"`        // 文件名称
	FileSha1         string    `gorm:"type:varchar(40);not null;unique;" json:"file_sha_1"` //文件哈希值
	FileSize         int64     `gorm:"type:varchar(20);not null;" json:"file_size"`         //文件大小
	Location         string    `gorm:"type:varchar(200);not null;" json:"location"`         //视频存放位置
	Extension        string    `gorm:"type:varchar(20);" json:"extension"`                  //扩展名
	Picture_Location string    `gorm:"type:varchar(200);not null;" json:"picture_location"` //视频图片存放位置
	Source           string    `gorm:"type:varchar(20);" json:"source"`                     //视频源
	CreatedAt        time.Time //视频创建时间
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
