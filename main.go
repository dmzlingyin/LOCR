package main

import (
	"fmt"
	"locr/cmd"
	C "locr/constant"
	"locr/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 是否输出到文件, 需结合管道符使用: cat test.jpg | locr -f > result.txt
	if C.IsSave {
		cmd.RecoPipe(os.Stdin)
		return
	}

	fmt.Println("locr start working...")
	go utils.Init()
	go cmd.Watch()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
