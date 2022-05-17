package utils

import (
	"fmt"
	"log"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func Init() {
	mainthread.Init(fn)
}

func fn() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	err := hk.Register()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("hotkey: %v is registered\n", hk)
	<-hk.Keydown()
	fmt.Printf("hotkey: %v is down\n", hk)
	<-hk.Keyup()
	fmt.Printf("hotkey: %v is up\n", hk)
	hk.Unregister()
	fmt.Printf("hotkey: %v is unregistered\n", hk)
}
