package commands

import "strings"

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

func (enum commandType) Command() string {
	switch enum {
	case CREATE:
		return "create_parking_lot"
	case PARK:
		return "park"
	case LEAVE:
		return "leave"
	case STATUS:
		return "status"
	case NUMBERS_WITH_COLOR:
		return "registration_numbers_for_cars_with_colour"
	case SLOT_WITH_NUMBER:
		return "slot_number_for_registration_number"
	case SLOTS_WITH_COLOR:
			return "slot_numbers_for_cars_with_colour"
	default:
		return ""
	}
}

type CommandBuilder struct{
	Executor string
	Parameters []string
}

func Parse(instruction string) *CommandBuilder{
	instructionSet := strings.Split(instruction, " ")
	command := CommandBuilder{
		Executor :  instructionSet[0],
		Parameters: instructionSet[1:],
	}
    return &command
}