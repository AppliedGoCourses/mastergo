package main

import "fmt"

func main() {

	var s string
	var n int

	fmt.Println("Please enter a word and a number. ")
	cnt, err := fmt.Scan(&s, &n) // & means we only pass a pointer to that variable.
	fmt.Println("You entered", cnt, "values:", s, "and", n)

	// Check the error
	if err != nil {
		fmt.Println("Scan failed:", err)
	}

	// Scanf
	fmt.Println("Please enter a word, press return, and enter and a number. ")
	cnt, err = fmt.Scanf("%s\n%d", &s, &n)

	if err != nil {
		fmt.Println("Scanf failed:", err)
	}

	fmt.Println("You entered", cnt, "values:", s, "and", n)
}
