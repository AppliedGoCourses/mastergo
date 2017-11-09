package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "abcde"
	for _, s := range s {
		s := unicode.ToUpper(s)
		fmt.Print(string(s))
	}
	fmt.Println("\n" + s)
}
