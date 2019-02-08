package commands

import "strings"

type command struct{
	Type string
	ParamCount int
}
type CommandBuilder struct{
	Executor string
	Parameters []string
}

type commandType string

const (
	CREATE             commandType = "create_parking_lot"
	PARK               commandType = "park"
	LEAVE              commandType = "leave"
	STATUS             commandType = "status"
	NUMBERS_WITH_COLOR  commandType = "registration_numbers_for_cars_with_colour"
	SLOT_WITH_NUMBER   commandType= "slot_number_for_registration_number"
	SLOTS_WITH_COLOR   commandType = "slot_numbers_for_cars_with_colour"
)

func (enum commandType) Get() *command {
	switch enum {
	case CREATE:
		return &command{"create_parking_lot",1}
	case PARK:
		return &command{"park",2}
	case LEAVE:
		return &command{"leave", 1}
	case STATUS:
		return &command{"status", 0}
	case NUMBERS_WITH_COLOR:
		return &command{"registration_numbers_for_cars_with_colour",1}
	case SLOT_WITH_NUMBER:
		return &command{"slot_number_for_registration_number",1}
	case SLOTS_WITH_COLOR:
			return &command{"slot_numbers_for_cars_with_colour",1}
	default:
		return &command{}
	}
}

func Parse(instruction string) *CommandBuilder{
	instructionSet := strings.Split(instruction, " ")
	command := CommandBuilder{
		Executor :  instructionSet[0],
		Parameters: instructionSet[1:],
	}
    return &command
}