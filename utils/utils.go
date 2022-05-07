package utils

import (
	"bytes"
	"locr/server"
	"strings"

	C "locr/constant"

	"golang.design/x/clipboard"
)

// IsImageFile 判断文件内容是否为图片类型(png/jpg/tif/webp)
func IsImageFile(content string) bool {
	if strings.HasPrefix(content, "file://") {
		switch content[len(content)-4:] {
		case ".png":
			return true
		case ".jpg":
			return true
		case ".tif":
			return true
		case ".bmp":
			return true
		default:
			return false
		}
	}
	return false
}

// IsImage 判断剪贴板内容是否为图片类型(png/jpg/tif/webp)
func IsImage(content []byte) bool {
	// 信息太少, 无法判断
	if len(content) < 10 {
		return false
	}

	// content 是否为png类型
	if bytes.Equal(content[:4], C.PNG[:4]) && bytes.Equal(content[len(content)-4:], C.PNG[4:]) {
		return true
	}
	// content 是否为jpg类型
	if bytes.Equal(content[:4], C.JPG[:4]) && bytes.Equal(content[len(content)-2:], C.JPG[4:]) {
		return true
	}
	// content 是否为bmp类型
	if bytes.Equal(content[:2], C.BMP) {
		return true
	}
	// content 是否为tiff类型
	if bytes.Equal(content[:4], C.TIFF) {
		return true
	}
	return false
}

// 如剪贴板无法使用，panic
func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

// ExtractText 提取识别结果的文本
func ExtractText(raw *server.Result) string {
	decReo := raw.Value[0]
	classify := strings.Split(decReo, "]]],")

	var res string
	for _, line := range classify {
		v := strings.Split(line, "'")[1]
		res += v
	}
	return res
}

// ExtractImage 将识别结果保存为图片
func ExtractImage(raw *server.Result) bool {
	return true
}
