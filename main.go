package main

import (
	"locr/cmd"
	"os"
	"os/signal"
	"syscall"

	C "locr/constant"
	_ "locr/pkg/hotkey"
	"locr/pkg/log"
)

func main() {
	// 是否输出到文件, 需结合管道符使用: cat test.jpg | locr -f > result.txt
	if C.IsSave {
		cmd.RecoPipe(os.Stdin)
		return
	}

	go cmd.Watch()
	log.InfoLogger.Println("locr start working...")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
