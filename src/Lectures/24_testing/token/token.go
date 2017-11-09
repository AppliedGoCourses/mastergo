package token

import "strings"

// Scan takes a string and splits it into words at white spaces (space, tab, newline).
func Scan(s string) []string {
	return strings.Split(s, " ") // wrong
	// return strings.Fields(s)  // correct
}
