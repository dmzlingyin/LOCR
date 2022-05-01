package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

const (
	VERSION = "0.1.0"
)

var (
	configPath string
	version    bool
)

func init() {
	flag.StringVar(&configPath, "p", "./config.yaml", "specify config file path")
	flag.BoolVar(&version, "v", false, "show current version of locr")
	flag.Parse()
	viper.SetConfigFile(configPath)
}

func main() {
	if len(flag.Args()) != 0 {
		flag.PrintDefaults()
	}

	if version {
		fmt.Println(VERSION)
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("no config founded.")
		} else {
			log.Fatal("config read failed.")
		}
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
