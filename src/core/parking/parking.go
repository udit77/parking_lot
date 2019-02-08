package parking

import (
	"github.com/parking_lot/src/core/database"
	"github.com/parking_lot/src/commands"
	"fmt"
)

type parkingModel struct{
	dbConnector *database.DBConnector
}


func New() *parkingModel{
	return &parkingModel{}
}

func (model *parkingModel) Init(){
	model.dbConnector = database.GetConnector()
}


func (model *parkingModel) Execute(instruction string){
	command := commands.Parse(instruction)
	if command.Executor == commands.CREATE.Command(){
		err := model.dbConnector.CreateLot(command)
		if err != nil{
			fmt.Println("error occurred while creating parking lot,please try again")
		}else{
			fmt.Println("Created a parking lot with",command.Parameters[0],"slots")
		}
		return
	}
	if command.Executor == commands.PARK.Command(){
		message , err := model.dbConnector.Park(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Println(message)
		}
		return
	}
	if command.Executor == commands.LEAVE.Command(){
		message , err := model.dbConnector.Leave(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Println(message)
		}
		return
	}

	if command.Executor == commands.STATUS.Command(){
		response , message, err := model.dbConnector.GetStatus(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Printf("%v%v%v\n","Slot No.  ","Registration No   ", "Colour")
			for i := range response{
				fmt.Printf("%v   %v   %v\n",response[i].Id,response[i].Number,response[i].Color)
			}
		}
		return
	}

	if command.Executor == commands.NUMBERS_WITH_COLOR.Command(){
		message , err := model.dbConnector.GetNumbersWithColor(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Println(message)
		}
		return
	}

	if command.Executor == commands.SLOT_WITH_NUMBER.Command(){
		message , err := model.dbConnector.GetSlotWithNumber(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Println(message)
		}
		return
	}

	if command.Executor == commands.SLOTS_WITH_COLOR.Command(){
		message , err := model.dbConnector.GetSlotsWithColor(command)
		if err != nil{
			if message != "" {
				fmt.Println(message)
				return
			}
			fmt.Println("error occurred while executing command please try again")
		}else{
			fmt.Println(message)
		}
		return
	}
}