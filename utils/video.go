package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"path/filepath"
)

// 设置视频文件所在的文件夹路径和保存图片的文件夹路径
const videoFolder = "./video"
const imageFolder = "images"

//想法 ：启动时将Deal_video加入到init函数中，之后一直监控文件夹

func Deal_video() {
	//遍历文件夹中的所有视频文件
	filepath.Walk(videoFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		//视频抽帧生成截图
		reader := ExampleReadFrameAsJpeg(path, 5)
		img, err := imaging.Decode(reader)
		if err != nil {
			fmt.Println(err)
		}

		err = imaging.Save(img, imageFolder+"/"+filepath.Base(path[:len(path)-4])+".jpg") //保存截图
		if err != nil {
			fmt.Println(err)
		}

		//todo 处理视频元信息

		return nil
	})

}

//todo 监控文件夹

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
