package utils

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"strconv"
	"strings"

	"locr/server"
)

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

	buff := new(bytes.Buffer)
	// png 支持任意image.Image类型 jpg.Encode只支持jpg类型
	err = png.Encode(buff, newImg)
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
