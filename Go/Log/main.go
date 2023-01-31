package main

import (
	"log"
	"os"
)

func createLog() {
	file, _ := os.Create("file.log")
	log.SetOutput(file)
	log.Println("Hello world")
	file.Close()
	log.SetOutput(os.Stdout)
	log.Println("Printing into standart out again")
}

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *aggregatedLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *aggregatedLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *aggregatedLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func main() {
	createLog()

	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "Info", flags)
	warnLogger := log.New(os.Stdout, "Warn", flags)
	errorLogger := log.New(os.Stdout, "Error", flags)
	infoLogger.Println("This is info log")
	warnLogger.Println("This is warn log")
	errorLogger.Println("This is error log")

	al := aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}

	al.info("This is info log")
	al.warn("This is warn log")
	al.error("This is error log")
}
