package main

import (
	"fmt"

	"github.com/appliedgocourses/bank"
)

func main() {

	err := bank.Load()
	if err != nil {
		fmt.Println("Cannot restore bank data.\n", err)
		return
	}
	defer func() {
		err := bank.Save()
		if err != nil {
			fmt.Println("Cannot save bank data.\n", err)
		}
	}()
}
