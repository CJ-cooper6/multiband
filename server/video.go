package server

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"multiband/dao"
	"multiband/model"
	"multiband/utils"
	"os"
	"path/filepath"
)

//想法 ：启动时将Deal_video加入到init函数中，之后一直监控文件夹

func InitDealVideo() {
	//遍历文件夹中的所有视频文件
	filepath.Walk(utils.VideoFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		//视频抽帧生成截图
		reader := ExampleReadFrameAsJpeg(path, 0)
		img, err := imaging.Decode(reader)
		if err != nil {
			fmt.Println(err)
		}
		err = imaging.Save(img, utils.ImageFolder+"/"+filepath.Base(path[:len(path)-4])+".jpg") //保存截图
		if err != nil {
			fmt.Println(err)
		}

		//处理视频元信息
		videometa := &model.VideoMeta{
			FileName:         filepath.Base(path),
			Location:         path,
			Extension:        filepath.Ext(path),
			Picture_Location: utils.ImageFolder + "/" + filepath.Base(path[:len(path)-4]) + ".jpg",
		}
		fi, err := os.Stat(path)
		if err != nil {
			return err
		}
		videometa.FileSize = fi.Size()

		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			fmt.Errorf("读取文件失败！")
			return err
		}

		videometa.FileSha1 = utils.FileSha1(file)
		fmt.Printf("%#v", videometa)
		dao.SaveVideoMeta(videometa)
		return nil
	})
}

//todo 监控文件夹

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg", "ss": 0, "t": 0.01}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println(err)
	}
	return buf
}
