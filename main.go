package main

import (
	"github.com/shimanekb/project1a/controller"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)

	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)

	args := os.Args
	if len(args) < 3 {
		log.Fatalln("Missing file path argument for input.")
	}

	filePath := args[1]
	outputPath := args[2]
	controller.ReadCsvCommands(filePath, outputPath)
}
