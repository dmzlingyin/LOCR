package utils

import (
	"log"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"

	C "locr/constant"
)

func Init() {
	mainthread.Init(fn)
}

// 热键：Ctrl + Shift + O, 控制自动识别开关
func fn() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyO)
	err := hk.Register()
	if err != nil {
		log.Println(err)
	}

	for {
		<-hk.Keydown()
		C.AutoReco = !C.AutoReco
	}
}
