package utility

import (
	"os"
	"github.com/pkg/errors"
)

var argsPassed []string

func ParseArgs() (bool, string, error){
	argsPassed = os.Args[1:]
	if len(argsPassed) == 0 {
		return false,"",nil
	}else {
		if len(argsPassed) == 1 {
			return true, argsPassed[0], nil
		}
		return false, "", errors.New("invalid input mode")
	}
}
