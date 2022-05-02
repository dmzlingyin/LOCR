package main

import (
	"fmt"
	"locr/cmd"
	_ "locr/constant"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("locr starting working...")
	go cmd.Watch()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
