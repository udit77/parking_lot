package parking

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"bufio"
	"io"
	"strings"
)

func (model *parkingModel) HandleCliInput(){
	for {
		prompt := promptui.Prompt{
			Label: "",
		}
		instruction, err := prompt.Run()
		if err != nil {
			log.Println("err : [Parking][HandleFileInput] error occurred in reading instruction, please try again")
			continue
		}
		model.Execute(instruction)
	}
}

func (model *parkingModel) HandleFileInput(filePath string){
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalln("fatal: [Parking][HandleFileInput] error occurred in reading input file", err)
	}

	reader := bufio.NewReader(file)
	for {
		instruction, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		instruction = strings.Trim(instruction,"\t\n")
		instruction = strings.TrimSpace(instruction)
		if err != nil {
			log.Fatalln("fatal: [Parking][HandleFileInput] error occurred in reading instruction from input file", err)
		}
		model.Execute(instruction)
	}
}