package main

import (
	"errors"
	"fmt"
	"os"
)

func helper(args []string) error {
	if len(args) > 1 {
		if args[1] == "help" {
			fmt.Println(`Usage:
<command_name> [help]`)
			os.Exit(0)
		}
		return errors.New("Unknown argument: " + args[1])
	}
	return nil
}
