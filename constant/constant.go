package constant

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	VERSION = "0.1.0"
)

var (
	// png文件头4个字节, 文件尾4个字节
	PNG = []byte{'\x89', '\x50', '\x4E', '\x47', '\xAE', '\x42', '\x60', '\x82'}
	// jpg文件头4个字节, 文件尾2个字节
	JPG = []byte{'\xFF', '\xD8', '\xFF', '\xE0', '\xFF', '\xD9'}
	// bmp只有文件头
	BMP = []byte{'\x42', '\x4D'}
	// tiff只有文件头
	TIFF = []byte{'\x49', '\x49', '\x2A', '\x00'}
)

var (
	configPath string
	IsSave     bool
	version    bool
	URL        string
	AutoReco   bool
)

var (
	// red, green, blue
	Colors = []color.RGBA{{255, 0, 0, 200}, {0, 255, 0, 200}, {0, 0, 255, 200}}
)

func init() {
	flag.StringVar(&configPath, "p", "./pconfig.yaml", "specify config file path")
	flag.BoolVar(&version, "v", false, "show current version of locr")
	flag.BoolVar(&IsSave, "f", false, "is save to file.")
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

	AutoReco = viper.GetBool("autoreco")
}
