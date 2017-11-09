package main

import "fmt"

func main() {

	// This code speaks for itself.
	// Just note the escape sequences used:
	// \t = tab
	// \n = newline
	// \" = "
	// \x61 = a

	simple := "A string literal."
	escape := "\tA string literal\nwith \"esc\x61pe sequences\"."

	raw := `I am
a raw string. 
  \n, \t etc. do not get evaluated.`

	fmt.Println("Simple literal:", simple)
	fmt.Println("With escape sequences:\n", escape)
	fmt.Println("Raw:", raw)

}
