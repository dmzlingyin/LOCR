package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path"
	"strings"
	"time"

	C "locr/constant"
)

const (
	PNG = iota
	JPG
	TIF
	BMP
)

// ImageFileType 判断文件内容是否为图片文件, 返回文件类型(png/jpg/tif/webp)
func ImageFileType(content string) int {
	if strings.HasPrefix(content, "file://") {
		switch content[len(content)-4:] {
		case ".png":
			return PNG
		case ".jpg":
			return JPG
		case ".tif":
			return TIF
		case ".bmp":
			return BMP
		default:
			return -1
		}
	}
	return -1
}

// ImageType 判断剪贴板内容是否为图片类型, 返回图片类型(png/jpg/tif/webp)
func ImageType(content []byte) int {
	// 信息太少, 无法判断
	if len(content) < 10 {
		return -1
	}

	// content 是否为png类型
	if bytes.Equal(content[:4], C.PNG[:4]) && bytes.Equal(content[len(content)-4:], C.PNG[4:]) {
		return PNG
	}
	// content 是否为jpg类型
	if bytes.Equal(content[:4], C.JPG[:4]) && bytes.Equal(content[len(content)-2:], C.JPG[4:]) {
		return JPG
	}
	// content 是否为tiff类型
	if bytes.Equal(content[:4], C.TIFF) {
		return TIF
	}
	// content 是否为bmp类型
	if bytes.Equal(content[:2], C.BMP) {
		return BMP
	}
	return -1
}

// 将识别结果保存到图片 存放路径：home目录
func saveResultToImage(imgByte []byte) error {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}

	homePath, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	name := time.Now().Format("2006-01-02 15:04:05")
	out, _ := os.Create(path.Join(homePath, name+".jpg"))
	fmt.Println(path.Join(homePath, name+".jpg"))
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100
	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		return err
	}

	return nil
}

// drawBox 绘制矩形框
func drawBox(img *image.RGBA, points [][]int) {
	for k, point := range points {
		rColor := C.Colors[k%len(C.Colors)]
		hline(img, point[0], point[1], point[2], rColor)
		hline(img, point[6], point[7], point[4], rColor)
		vline(img, point[1], point[0], point[7], rColor)
		vline(img, point[3], point[2], point[5], rColor)
	}
}

// hline 画横线
func hline(img *image.RGBA, x1, y, x2 int, col color.Color) {
	for x1 <= x2 {
		img.Set(x1, y, col)
		x1++
	}
}

// vline 画竖线
func vline(img *image.RGBA, y1, x, y2 int, col color.Color) {
	for y1 <= y2 {
		img.Set(x, y1, col)
		y1++
	}
}
