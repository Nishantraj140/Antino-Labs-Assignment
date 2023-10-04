package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DBLogger      *log.Logger
	File          *os.File
)

func InitLogger(fileName string) {
	File, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(File, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(File, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(File, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DBLogger = log.New(File, "DBLOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
