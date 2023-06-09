package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"multiband/utils"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func IntDb() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		DisableForeignKeyConstraintWhenMigrating: true, //禁止自动创建外键约束
	})
	if err != nil {
		fmt.Println("连接数据库出错：", err)
	}

	_ = Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8;").AutoMigrate(&VideoMeta{})

	sqlDB, _ := Db.DB()
	//设置连接池u中的最大闲置连接数
	sqlDB.SetMaxIdleConns(10)

	//设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(100)

	//设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
