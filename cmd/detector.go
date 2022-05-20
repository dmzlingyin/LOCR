package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.design/x/clipboard"

	C "locr/constant"
	"locr/pkg/utils"
	"locr/server"
)

type Detector interface {
	Detect()
	Recognition()
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

		if C.AutoReco {
			img.Recognition()
		}
		cancel()
	}
}

func (img *ImageDetector) Recognition() {
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
			err = utils.ExtractImage(content, res)
			if err != nil {
				log.Println(err)
			}
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

		if C.AutoReco {
			shot.Recognition()
		}
		cancel()
	}
}

func (shot *ShotDetector) Recognition() {
	if utils.IsImage(shot.Data) {
		res, err := server.RecoBase64(shot.Data)
		if err != nil {
			log.Println(err)
		} else {
			shot.Result = utils.ExtractText(res)
			err = utils.ExtractImage(shot.Data, res)
			if err != nil {
				log.Println(err)
			}
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
