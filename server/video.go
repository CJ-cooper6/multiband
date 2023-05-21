package server

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fsnotify/fsnotify"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"multiband/dao"
	"multiband/model"
	"multiband/utils"
	"os"
	"path/filepath"
	"time"
)

func InitDealVideo() {
	go DealVideo(utils.VideoFolder1, utils.ImageFolder1)
}

//生成视频元信息并存储
func DealVideo(VideoFolder string, ImageFolder string) {
	//遍历文件夹中的所有视频文件
	filepath.Walk(VideoFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		return CreateVideoInfo(path, VideoFolder, ImageFolder)
	})

	//监控文件夹
	// 创建一个新的文件系统通知监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = filepath.Walk(VideoFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 将要监听的文件夹添加到监视器
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 启动一个goroutine来处理文件系统事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 处理文件的创建事件
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("New file created:", event.Name)
					time.Sleep(5 * time.Second)
					_ = CreateVideoInfo(event.Name, VideoFolder, ImageFolder)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				// 处理错误
				log.Println("Error:", err)
			}
		}
	}()

	// 主goroutine阻塞，持续监听文件夹变化
	select {}
}

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

func CreateVideoInfo(path, VideoFolder, ImageFolder string) error {
	//视频抽帧生成截图
	reader := ExampleReadFrameAsJpeg(path, 0)
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}
	err = imaging.Save(img, ImageFolder+"/"+filepath.Base(path[:len(path)-4])+".jpg") //保存截图
	if err != nil {
		fmt.Println(err)
	}

	//处理视频元信息
	videometa := &model.VideoMeta{
		FileName:         filepath.Base(path),
		Location:         path,
		Extension:        filepath.Ext(path),
		Picture_Location: ImageFolder + "/" + filepath.Base(path[:len(path)-4]) + ".jpg",
		Source:           filepath.Base(VideoFolder),
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
}
