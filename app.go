package main

import (
	"github.com/parking_lot/src/core/parking"
	"github.com/manifoldco/promptui"
	"log"
)


func main() {
	parkingModel := parking.New()
	parkingModel.Init()
	for {
		prompt := promptui.Prompt{
			Label:"",
		}
		command , err := prompt.Run()
		if err != nil {
			log.Println("err : [Main] error occurred in parsing command, please try again")
			continue
		}
		parkingModel.Execute(command)
	}
}