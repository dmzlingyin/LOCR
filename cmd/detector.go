package cmd

import (
	"context"
	"fmt"
	"locr/server"
	"log"
	"os"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func Watch() {
	for {
		// 24小时监听剪贴板变化
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)
		changed := clipboard.Watch(ctx, clipboard.FmtText)
		content := string(<-changed)

		fmt.Println(content)
		if isImageFile(content) {
			res, err := Recognation(content)
			if err != nil {
				log.Println(err)
			} else {
				// 识别结果写入剪贴板, 并等待写入成功
				done := clipboard.Write(clipboard.FmtText, []byte(res))
				<-done
			}
		}
		cancel()
	}
}

// isImageFile 判断剪贴板内容是否为图片类型(png/jpg/tif/webp)
func isImageFile(content string) bool {
	if strings.HasPrefix(content, "/") {
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

// Recognation 对图片进行文字识别, 如果成功返回识别内容, 否则返回错误信息
func Recognation(content string) (string, error) {
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
