package log

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("log.txt open fail.")
	}

	InfoLogger = log.New(file, "info: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "error: ", log.Ldate|log.Ltime|log.Lshortfile)
}
