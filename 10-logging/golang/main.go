package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logWithBuiltInLoggerToConsole()
	logWithBuiltInLoggerToFile()
	logWithZap()
}

func logWithZap() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	logger.Info("Hey, I'm using Zap!")
}

func logWithBuiltInLoggerToConsole() {
	log.SetPrefix("main():logWithBuiltInLoggerToConsole():")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Hey, I'm using Console!")
	//log.Fatal("Hey, I'm using Fatal!") // ends program
	//log.Panic("Hey, I'm using Panic!") // ends program
}

func logWithBuiltInLoggerToFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("main():logWithBuiltInLoggerToFile():")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("HHey, I'm using File!")
}
