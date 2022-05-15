package main

import (
	"fmt"
	"locr/cmd"
	C "locr/constant"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 是否输出到文件, 需结合管道符使用: cat test.jpg | locr -f > result.txt
	if C.IsSave {
		cmd.RecoPipe(os.Stdin)
		os.Exit(0)
	}

	fmt.Println("locr start working...")
	go cmd.Watch()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
