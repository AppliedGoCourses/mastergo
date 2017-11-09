package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Restore the bank data.
	// Call bank.Load() here and handle any error.

	// TODO: your code here

	// Save the bank data.
	// Use a deferred function for calling bank.Save().
	// (See the lecture on function behavior.)
	// If Save() returns an error, print it.

	// TODO: your code here

	// Perform the action.
	// os.Args[0] is the path to the executable.
	// os.Args[1] is the first parameter - the action we want to perform:
	// create, list, update, transfer, or history.

	switch os.Args[1] {

	case "list":

		// TODO: case code here

	// TODO: more cases

	default:

		// TODO: your code here

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
}

// update takes a name and an amount, deposits the amount if it
// is greater than zero, or withdraws it if it is less than zero,
// and returns the new balance and any error that occurred.
func update(name string, amount int) (int, error) {

	// TODO: your code here

}

// transfer takes two names and an amount, transfers the amount from
// the account belonging to name #1 to the account belonging to name #2,
// and returns the new balances of both accounts and any error that occurred.
func transfer(name, name2 string, amount int) (int, int, error) {

	// TODO: your code here

}

// history takes an account name, retrieves the account, and calls bank.History()
// to get the history closure function. Then it calls the closure in a loop,
// formatting the return values and appending the result to the output string, until the boolean return parameter of the closure is `false`.
func history(name string) (string, error) {

	// TODO: your code here

}
