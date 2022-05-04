package utils

import (
	"locr/server"
	"strings"
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
func ExtractImage() {

}
