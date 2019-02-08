package utility

import "os"

var (
	arguementsPassed []string
)


func ParseArgs() bool{
	arguementsPassed = os.Args[1:]
	if len(arguementsPassed) == 0 {

	}else {
		if len(arguementsPassed) == 1 {

		}
	}
}
