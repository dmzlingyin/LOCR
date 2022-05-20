package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"locr/server"
	"os"
	"path"
	"strconv"
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
func ExtractImage(data []byte, raw *server.Result) error {
	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	drawBox(dst, extractPoints(raw))
	newImg := dst.SubImage(img.Bounds())

	var opts jpeg.Options
	opts.Quality = 100
	buff := new(bytes.Buffer)
	err = jpeg.Encode(buff, newImg, &opts)
	if err != nil {
		return err
	}

	return saveResultToImage(buff.Bytes())
}

// extractPoints 提取识别结果的坐标点
// 例如: [28 18 155 15 155 29 28 32], 按顺序分组(28, 18), (155, 15), (155, 29), (28, 32), 分别代表检测结果的四个定位点: 左上, 右上, 右下, 左下
func extractPoints(raw *server.Result) [][]int {
	decReo := raw.Value[0]
	splited := strings.Split(decReo, "]], [")
	res := make([][]int, len(splited))

	for k, item := range splited {
		res[k] = make([]int, 8)
		points := strings.Split(item, "), [")[1]
		lists := strings.Split(points, ", ")

		for i, v := range lists {
			if i%2 == 0 {
				a, _ := strconv.ParseFloat(v[1:], 64)
				res[k][i] = int(a)
			} else {
				b, _ := strconv.ParseFloat(v[:len(v)-1], 64)
				res[k][i] = int(b)
			}
		}
	}
	return res
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
	out, _ := os.Create(path.Join(homePath, "test.jpg"))
	fmt.Println(path.Join(homePath, "test.jpg"))
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
