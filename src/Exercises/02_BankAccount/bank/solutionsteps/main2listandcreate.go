package main

import (
	"fmt"
	"os"

	"github.com/appliedgocourses/bank"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Restore the bank data.
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

	// Perform the action.
	switch os.Args[1] {

	case "list":

		fmt.Println(bank.ListAccounts())

	case "create":
		if len(os.Args) < 3 {
			usage()
		}

		name := os.Args[2]
		if a, err := bank.GetAccount(name); err == nil && a != nil {
			fmt.Println("Account '", name, "' already exists.")
			return
		}
		bank.NewAccount(name)
		fmt.Println("Account '", name, "' created.")

	default:

		fmt.Println("Unknown command:", os.Args[1])
		usage()

	}
}

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
	os.Exit(1) // Deferred functions are NOT called if exiting this way!
}
