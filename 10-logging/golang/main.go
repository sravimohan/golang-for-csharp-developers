package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func main() {
	logToConsole()
	logToFile()
	logWithZap()
}

func logToConsole() {
	log.SetPrefix("logToConsole() : ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Hello World 1!")
	// log.Fatal("I am Fatal") // program will end with this line
	// log.Panic("I am Panic") // program will end with this line with additional context
	log.Println("Hello World 2!")
}

func logToFile() {
	file, err := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("cannot open log file")
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetPrefix("logToFile() : ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Hello World 1!")
	log.Println("Hello World 2!")
}

func logWithZap() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("cannot init Zap")
	}

	defer logger.Sync()
	logger.Info("Info from Zap")
}
