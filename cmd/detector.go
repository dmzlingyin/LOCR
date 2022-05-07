package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"locr/server"
	"locr/utils"
	"log"
	"os"
	"time"

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
	if utils.IsImageFile(img.Data) {
		reader, err := os.Open(img.Data[7:])
		if err != nil {
			log.Println(err)
		}
		defer reader.Close()

		content, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Println(err)
		}

		res, err := server.RecoBase64(content)
		if err != nil {
			log.Println(err)
		} else {
			img.Result = utils.ExtractText(res)
			clipboard.Write(clipboard.FmtText, []byte(img.Result))
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
	if utils.IsImage(shot.Data) {
		res, err := server.RecoBase64(shot.Data)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(res)
			shot.Result = utils.ExtractText(res)
			clipboard.Write(clipboard.FmtText, []byte(shot.Result))
		}
	}
}

func Watch() {
	imageDetector := &ImageDetector{}
	shotDetector := &ShotDetector{}

	go imageDetector.Detect()
	go shotDetector.Detect()
}
