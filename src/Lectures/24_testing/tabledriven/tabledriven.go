package tabledriven

import (
	"errors"
	"strings"
)

// Scan takes a string and splits it into words at white spaces (space, tab, newline).
// If the input string is empty, Scan returns an error.
func Scan(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("s must not be empty")
	}
	return strings.Fields(s), nil
}
