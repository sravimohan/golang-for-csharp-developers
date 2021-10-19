package main

import (
	"log"
	"os"
)

func main() {
	logToFile()

	log.SetPrefix("main(): ")
	log.Print("Hey, I'm using Print!")
	//log.Fatal("Hey, I'm using Fatal!") // ends program
	//log.Panic("Hey, I'm using Panic!") // ends program
	log.Print("Program End!")
}

func logToFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
}
