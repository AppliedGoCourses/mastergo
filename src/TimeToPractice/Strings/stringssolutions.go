package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {
	afterSpace := false
	for i, c := range s {
		// If the first character of the string (i == 0) is a letter,
		// or if the current character is a letter and the previous was
		// a space, then ensure the letter is uppercase and append it
		// to the output string.
		if (afterSpace || i == 0) && unicode.IsLetter(c) {
			// c is a rune, hence we need uincode.ToUpper().
			acr += string(unicode.ToUpper(c))
			afterSpace = false
		}
		// again, c is a rune (int32), hence `c == " "` does not work here.
		if unicode.IsSpace(c) {
			afterSpace = true
		}
	}
	return acr
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
