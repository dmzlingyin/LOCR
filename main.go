package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	flag.StringVar(&configPath, "config", "./config.yaml", "please specify config pagh")
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()
	viper.SetConfigFile(configPath)

}

func main() {
	if len(flag.Args()) != 0 {
		flag.PrintDefaults()
	}

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("no config founded.")
		} else {
			log.Fatal("config read failed.")
		}
		os.Exit(1)
	}
}
