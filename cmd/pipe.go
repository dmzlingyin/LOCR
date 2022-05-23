package cmd

import (
	"fmt"
	"io"

	"locr/pkg/log"
	"locr/pkg/utils"
	"locr/server"
)

// RecoPipe 适用于一次性识别图片内容(无剪贴板), 结合管道符, 将识别结果输出到文件
func RecoPipe(r io.Reader) {
	b, err := io.ReadAll(r)
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	if utils.ImageType(b) == -1 {
		log.WarningLogger.Println("File is not an image.")
		return
	}

	res, err := server.RecoBase64(b)
	if err != nil {
		log.ErrorLogger.Println(err)
	}

	fmt.Println(utils.ExtractText(res))
}
