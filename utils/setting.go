package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	VideoFolder1 string // 设置视频文件所在的文件夹路径和保存图片的文件夹路径
	ImageFolder1 string
	VideoFolder2 string
	ImageFolder2 string
	Trans        string
)

func Setting() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，检查文件路径：", err)
	}

	LoadServer(file)
	LoadDb(file)
	LoadVideo(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("Dbport").MustString("3306")
	DbUser = file.Section("database").Key("Dbuser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("ACAT")
}

func LoadVideo(file *ini.File) {
	VideoFolder1 = file.Section("video").Key("VideoFolder1").MustString("")
	ImageFolder1 = file.Section("video").Key("ImageFolder1").MustString("")
	VideoFolder2 = file.Section("video").Key("VideoFolder2").MustString("")
	ImageFolder2 = file.Section("video").Key("ImageFolder2").MustString("")
	Trans = file.Section("video").Key("Trans").MustString("")
}
