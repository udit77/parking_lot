package commands

import (
	"testing"
	"strings"
	"fmt"
)

func TestParse(t *testing.T) {
	instructionSet := strings.Split("park KA-01-HH-1234 White", " ")
	command := CommandBuilder{
		Executor :  instructionSet[0],
		Parameters: instructionSet[1:],
	}
	fmt.Println(command)
}
