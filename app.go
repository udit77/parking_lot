package main

import (
	"github.com/parking_lot/src/core/parking"
	"log"
	"github.com/parking_lot/src/utility"
	"path/filepath"
	"os"
)


func main() {
	_, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("fatal : [Main] error in resolution of file path",err)
	}

	isFileMode , inputFile, err := utility.ParseArgs()
	if err != nil {
		log.Fatalln("fatal : [Main] invalid input mode")
	}

	parkingModel := parking.New()
	parkingModel.Init()

	if !isFileMode {
		parkingModel.HandleCliInput()
	}else{
		_, err := os.Stat(inputFile)
		if err != nil {
			log.Fatalln("fatal: [Main] error occurred in reading input file", err)
		} else {
			parkingModel.HandleFileInput(inputFile)
		}
	}
}