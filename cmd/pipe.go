package cmd

import (
	"fmt"
	"io"
	"locr/server"
	"locr/utils"
	"log"
)

// RecoPipe 适用于一次性识别图片内容(无剪贴板), 结合管道符, 将识别结果输出的文件
func RecoPipe(r io.Reader) {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Println(err)
	}

	res, err := server.RecoBase64(b)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(utils.ExtractText(res))
}
