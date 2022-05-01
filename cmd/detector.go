package cmd

import (
	"fmt"

	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func Detector() string {
	text := clipboard.Read(clipboard.FmtText)
	fmt.Println(string(text))
	return ""
}
