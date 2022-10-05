package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Solution 1
//
// A straightforward approach is to loop over the string, test if
// (a) we're just one letter after a space, and
// (b) the current character is a letter,
// and then turn this letter into uppercase.

func acronymWithLoop(s string) (acr string) {
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

// Optional task: optimize the loop
//
// Concatenating strings with + or += has some hidden cost. 
// Strings are immutable. Concatenating two strings creates a new string,
// and the original strings are subject to garbage collection. 
// Hence concatenating strings in a loop may put some load on the garbage collector
// if the loop is large enough.
//
// Your optional task: optimize the string concatenation.
//
// Have a look at the Builder type in package strings. The String Builder provides 
// efficient string manipulation methods that minimize copying and heap allocations.
// Re-write the loop to replace the += operator with String Builder methods.



// Solution 2
//
// (Author: Adil Billa)
// Exploring the `strings` package you will find a handy function called
// "Map()" that maps a function to every rune of a given string,
// thus hiding the loop behind a simple function call.
// Furthermore, to save looking for a preceding space, we can simply
// turn all first letters of each word to uppercase and then just
// check the letter's case in order to find the first letter in a word.

func acronymWithMap(s string) string {
	extractUpperCase := func(r rune) rune {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			return r
		}
		return -1
	}

	return strings.Map(extractUpperCase, strings.Title(s))
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronymWithLoop(s))
	fmt.Println(acronymWithMap(s))
}
