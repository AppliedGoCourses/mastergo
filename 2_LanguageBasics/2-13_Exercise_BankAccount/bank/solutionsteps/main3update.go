package main

import (
	"fmt"
	"os"
	"strconv"

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

	case "update":

		if len(os.Args) < 4 {
			usage()
		}
		name := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(os.Args[3], "is not a valid integer")
			return
		}

		bal, err := update(name, amount)
		if err != nil {
			fmt.Println(errors.WithStack(err))
			return
		}

		lastaction := "Deposited"
		if amount < 0 {
			lastaction = "Withdrawn"
		}
		fmt.Printf("Acount '%s': %s %d credits.\nNew balance: %d\n", name, lastaction, amount, bal)

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
	os.Exit(1)
}

func update(name string, amount int) (int, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return 0, errors.Wrap(err, "account not found")
	}
	if amount == 0 {
		return bank.Balance(account), errors.New("amount must not be zero")
	}

	balance := 0
	if amount > 0 {
		balance, err = bank.Deposit(account, amount)
		if err != nil {
			return balance, errors.Wrap(err, "depositing failed")
		}
	} else { // amount < 0
		// Note: we must negate the amount here. bank.Withdraw() expects a positive value.
		balance, err = bank.Withdraw(account, -amount)
		if err != nil {
			return balance, errors.Wrap(err, "withdrawing failed")
		}
	}
	return balance, nil
}
