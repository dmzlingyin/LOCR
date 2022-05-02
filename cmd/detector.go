package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"locr/server"

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
	content := string(text)[7:]
	if strings.HasSuffix(content, ".png") {
		file, err := os.Open(content)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		res, err := server.RecoFile(file)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(res.Result)
		fmt.Println([]byte(res.Result))
		clipboard.Write(clipboard.FmtText, []byte(res.Result))
		// <-ch
		fmt.Println("clipboard write success.")
	}
	return ""
}
