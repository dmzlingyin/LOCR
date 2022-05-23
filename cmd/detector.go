package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.design/x/clipboard"

	C "locr/constant"
	"locr/pkg/log"
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
	if utils.ImageFileType(img.Data) != -1 {
		reader, err := os.Open(img.Data[7:])
		if err != nil {
			log.ErrorLogger.Println(err)
		}
		defer reader.Close()

		content, err := ioutil.ReadAll(reader)
		if err != nil {
			log.ErrorLogger.Println(err)
		}

		res, err := server.RecoBase64(content)
		if err != nil {
			log.ErrorLogger.Println(err)
		} else {
			img.Result = utils.ExtractText(res)
			clipboard.Write(clipboard.FmtText, []byte(img.Result))

			// 将检测结果保存为图片
			go func() {
				err = utils.ExtractImage(content, res)
				if err != nil {
					log.ErrorLogger.Println(err)
				}
			}()
		}
	} else {
		log.InfoLogger.Println("new content of clipboard is not a image file.")
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
	if utils.ImageType(shot.Data) != -1 {
		res, err := server.RecoBase64(shot.Data)
		if err != nil {
			log.ErrorLogger.Println(err)
		} else {
			shot.Result = utils.ExtractText(res)
			clipboard.Write(clipboard.FmtText, []byte(shot.Result))

			go func() {
				err = utils.ExtractImage(shot.Data, res)
				if err != nil {
					log.ErrorLogger.Println(err)
				}
			}()
		}
	} else {
		log.InfoLogger.Println("new content of clipboard is not a shot image.")
	}
}

// 如剪贴板无法使用, 退出程序
func init() {
	err := clipboard.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Watch() {
	imageDetector := &ImageDetector{}
	shotDetector := &ShotDetector{}

	go imageDetector.Detect()
	go shotDetector.Detect()
}
