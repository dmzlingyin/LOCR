package constant

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type ocrServer struct {
	Addr string
	Port string
}

const (
	VERSION = "0.1.0"
)

var (
	configPath string
	version    bool
	OcrServer  *ocrServer
)

func Init() {
	flag.StringVar(&configPath, "p", "./pconfig.yaml", "specify config file path")
	flag.BoolVar(&version, "v", false, "show current version of locr")
	flag.Parse()
	viper.SetConfigFile(configPath)

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
	OcrServer = NewOcrServer()
}

func NewOcrServer() *ocrServer {
	return &ocrServer{
		Addr: viper.GetStringMapString("ocrserver")["address"],
		Port: viper.GetStringMapString("ocrserver")["port"],
	}
}
