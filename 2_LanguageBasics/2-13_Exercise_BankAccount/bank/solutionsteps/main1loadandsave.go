package main

import (
	"fmt"

	"github.com/appliedgocourses/bank"
	"github.com/pkg/errors"
)

func main() {

	err := bank.Load()
	if err != nil {
		fmt.Println("Cannot restore bank data.\n", errors.WithStack(err))
		return
	}
	defer func() {
		err := bank.Save()
		if err != nil {
			fmt.Println("Cannot save bank data.\n", errors.WithStack(err))
		}
	}()
}
