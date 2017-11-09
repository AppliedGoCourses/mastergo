package procinfo

import (
	"os"
	"strconv"
)

func Get() string {
	str := "Process info:\n"
	str += "PID: " + strconv.Itoa(os.Getpid()) + "\n"
	wd, err := os.Getwd()
	if err != nil {
		wd = err.Error()
	}
	str += "Working dir: " + wd
	return str
}
