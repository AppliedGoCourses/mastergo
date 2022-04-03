package fuzzing

import (
	"strings"
)

// Scan takes a string and splits it into words at white spaces
// (space, tab, newline).
// If the input string is empty, Scan returns an empty slice.
// func Scan(s string) []string {
// 	if s == "" {
// 		return []string{}
// 	}
// 	return strings.Split(s, " ")
// }

// Iterative improvements of Scan(). Comment the original one, and uncomment each of the following versions to repeat the problem-solving journey from the lecture.

// func Scan(s string) []string {
// 	trimmed := strings.TrimSpace(s)
// 	if trimmed == "" {
// 		return []string{}
// 	}
// 	return strings.Split(trimmed, " ")
// }

func Scan(s string) []string {
	return strings.Fields(s)
}
