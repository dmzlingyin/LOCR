package cmd

import (
	"bytes"
	"context"
	"fmt"
	"locr/server"
	"log"
	"os"
	"strings"
	"time"

	C "locr/constant"

	"golang.design/x/clipboard"
)

type Detector interface {
	Detect()
	Recognation()
}

// 图片文件(存放在磁盘)
type ImageDetector struct {
	Data   string
	Result string
}

// 截图(存放在内存)
type ShotDetector struct {
	Data   []byte
	Result string
}

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func (img *ImageDetector) Detect() {
	for {
		// 24小时监听剪贴板变化
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
		text := clipboard.Watch(ctx, clipboard.FmtText)
		img.Data = string(<-text)
		img.Recognation()

		cancel()
	}
}

func (img *ImageDetector) Recognation() {
	if isImageFile(img.Data) {
		res, err := imageRecognation(img.Data[7:])
		if err != nil {
			log.Println(err)
		} else {
			img.Result = res
			// 识别结果写入剪贴板, 并等待写入成功
			done := clipboard.Write(clipboard.FmtText, []byte(res))
			fmt.Println(img.Result)
			<-done
		}
	}
}

func (shot *ShotDetector) Detect() {
	for {
		// 24小时监听剪贴板变化
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
		img := clipboard.Watch(ctx, clipboard.FmtImage)
		shot.Data = <-img
		shot.Recognation()

		cancel()
	}
}

func (shot *ShotDetector) Recognation() {
	if isImage(shot.Data) {
		res, err := shotRecognation(shot.Data)
		if err != nil {
			log.Println(err)
		} else {
			shot.Result = res
			// 识别结果写入剪贴板, 并等待写入成功
			done := clipboard.Write(clipboard.FmtText, []byte(res))
			fmt.Println(shot.Result)
			<-done
		}
	}
}

func Watch() {
	imageDetector := &ImageDetector{}
	shotDetector := &ShotDetector{}

	go imageDetector.Detect()
	go shotDetector.Detect()
}

// isImageFile 判断文件内容是否为图片类型(png/jpg/tif/webp)
func isImageFile(content string) bool {
	if strings.HasPrefix(content, "file://") {
		switch content[len(content)-4:] {
		case ".png":
			return true
		case ".jpg":
			return true
		case ".tif":
			return true
		case ".webp":
			return true
		default:
			return false
		}
	}
	return false
}

// isImage 判断剪贴板内容是否为图片类型(png/jpg/tif/webp)
func isImage(content []byte) bool {
	if len(content) < 10 {
		return false
	}
	magic := content[:8]
	if bytes.Equal(magic, C.PNG) {
		return true
	}
	return false
}

// imageRecognation 对图片进行文字识别, 如果成功返回识别内容, 否则返回错误信息
func imageRecognation(content string) (string, error) {
	// 打开待识别图片
	file, err := os.Open(content)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 图片识别
	res, err := server.RecoFile(file)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// shotRecognation 对截图进行文字识别, 如果成功返回识别内容, 否则返回错误信息
func shotRecognation(content []byte) (string, error) {
	res, err := server.RecoBase64(content)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}
