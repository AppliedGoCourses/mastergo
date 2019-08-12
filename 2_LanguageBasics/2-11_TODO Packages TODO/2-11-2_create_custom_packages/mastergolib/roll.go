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

// seed starts with a lowercase letter and hence is not exported.
func seed() {
	// time.Now() is sufficiently arbitrary to serve as a
	// random seed.
	// UnixNano() converts a time.Time value into an int64
	// type that Seed() expects.
	rand.Seed(time.Now().UnixNano())
}

// init is a special function. Each package (including main) can
// have an init function. When a binary starts, the init functions
// of all packages are run before entering main().
// This is a god place to initialize a library package.
// Here, we seed the random generator so that it does not
// generate the same series of numbers every time the binary runs.
func init() {
	seed()
}
