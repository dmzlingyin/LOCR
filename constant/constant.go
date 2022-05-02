package constant

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
	PNG        = []byte{137, 80, 78, 71, 13, 10, 26, 10}
	configPath string
	version    bool
	URL        string
)

func init() {
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
	addr := viper.GetStringMapString("ocrserver")["address"]
	port := viper.GetStringMapString("ocrserver")["port"]
	URL = addr + ":" + port + "/"
}
