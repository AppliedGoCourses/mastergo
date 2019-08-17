// Package dice lets you roll the dice.
package dice

import (
	"math/rand"
	"time"
)

// Roll rolls a dice with n sides.
func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

func Seed(n int64) {
	// For n == 0, create a random seed value.
	if n == 0 {
		// time.Now() is sufficiently arbitrary to serve as a
		// random seed.
		// UnixNano() converts a time.Time value into an int64
		// type that Seed() expects.
		n = time.Now().UnixNano()
	}
	// Otherwise, use n as seed value.
	rand.Seed(n)
}

// init is a special function. Each package (including main) can
// have an init function. When a binary starts, the init functions
// of all packages are run before entering main().
//
// However, there are two good reasons for not using init():
// 1. init() exists mainly for initializing the global state of a
//    package, and global package state is generally a bad idea.
// 2. If there are more than one init() function in your code, they
//    are executed in an undefined sequence, hence the result is
//    not deterministic and thus can lead to subtle bugs.
//
// func init() {
// 	seed()
// }
