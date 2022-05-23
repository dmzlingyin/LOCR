package hotkey

import (
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"

	C "locr/constant"
	"locr/pkg/log"
)

func init() {
	go mainthread.Init(fn)
}

// 热键：Ctrl + Shift + O, 控制自动识别开关
func fn() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyO)
	err := hk.Register()
	if err != nil {
		log.ErrorLogger.Println(err)
	}

	for {
		<-hk.Keydown()
		C.AutoReco = !C.AutoReco
		if C.AutoReco {
			log.InfoLogger.Println("autoreco on")
		} else {
			log.InfoLogger.Println("autoreco off")
		}
	}
}
