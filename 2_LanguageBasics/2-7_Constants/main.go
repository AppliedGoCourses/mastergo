package main

import (
	"fmt"
	"math"
)

func main() {
	// Basics

	const kilo = 1024

	const rock, paper, scissors = 1, 2, 3

	const (
		cvalue = 299792458
		cunit  = "m/s"
		cname  = "speed of light"
	)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// Literals

	// Note that fmt.Println() may change the format.

	// Runes
	fmt.Println("\n*** RUNES ***")

	fmt.Println('😀', string('😀'))
	fmt.Println('\377', string('\377'))
	fmt.Println('\u266B', string('\u266B'))
	fmt.Println('\U0001F600', string('\U0001F600'))

	// Integers
	fmt.Println("\n*** INTEGERS ***")

	fmt.Println(1337)
	fmt.Println(01337)
	fmt.Println(0x1337)
	fmt.Println(0x1CAFE42)

	// Floats
	fmt.Println("\n*** FLOATS ***")

	fmt.Println(1e5)
	fmt.Println(1e-3)
	fmt.Println(123e20)
	fmt.Println(1.234)
	fmt.Println(1.234e9)
	fmt.Println(2.71828)
	fmt.Println(2.)
	fmt.Println(.71828)
	fmt.Println(2.e+7)
	fmt.Println(.5e-10)

	// Imaginaries
	fmt.Println("\n*** IMAGINARY ***")

	fmt.Println(1e5i)
	fmt.Println(1e-3i)
	fmt.Println(123e20i)
	fmt.Println(1.234i)
	fmt.Println(1.234e9i)
	fmt.Println(2.71828i)
	fmt.Println(2.i)
	fmt.Println(.71828i)
	fmt.Println(2.e+7i)
	fmt.Println(.5e-10i)

	// Strings
	fmt.Println("\n*** STRINGS ***")

	fmt.Println("A string 😀")
	fmt.Println("A\tstring")
	fmt.Println("A string \U0001F600")

	// Booleans
	fmt.Println("\n*** BOOLEANS ***")

	fmt.Println(true)
	fmt.Println(false)

	fmt.Println()

	// Untyped constants
	fmt.Println("\n*** UNTYPED CONSTANTS ***")

	const x = 3.0 + 0i

	var i32 int32 = x
	var i64 int64 = x
	var f32 float32 = x
	var f64 float64 = x
	var c64 complex64 = x
	var c128 complex128 = x
	fmt.Println("Untyped complex literal:", i32, i64, f32, f64, c64, c128)

	// const x = 3.1 + 0i // fails getting assigned to int32, int64
	// const x = 3 + 7i // fails getting assigned to all but the complex types

	// Typed constants

	const pi32 float32 = math.Pi

	// Arbitrary precision for integers and floats
	fmt.Println("\n*** ARBITRARY PRECISION ***")

	const pi = 3.14159265358979323846264338327950288419716939937510582097494459 // math.Pi
	const pi64 float64 = math.Pi

	// A literal cannot be printed with its full precision - it is turned into a float type before. Compare these two outputs:
	fmt.Println("Untyped: ", pi)
	fmt.Println("Typed: ", pi64)

	const huge = 1 << 100 // 2^100 does not fit into any integer
	const notSoHuge = 1 << 90
	const smallEnough = huge / notSoHuge // 2^10 = 1024

	// large := huge  // constant 1267650600228229401496703205376 overflows int
	smallish := smallEnough
	fmt.Println("Small enough:", smallish)

	// Value repetition
	fmt.Println("\n*** VALUE REPETITION ***")

	const (
		twelve = 12 // 12
		dozen       // 12
		months      // 12
	)

	fmt.Println("Repeating value: ", twelve, dozen, months)

	// iota: generating constants

	// Enumerations
	fmt.Println("\n*** IOTA ENUMERATIONS ***")

	const (
		zero  = iota // 0
		one          // 1
		two          // 2
		three = iota // 3, iota has no effect here
		four         // 4
		five         // 5
	)

	const (
		ten    = iota*10 + 10 // 10
		twenty                // 20
		thirty                // 30
	)

	fmt.Println("iota:", zero, one, two, three, four, five, ten, twenty, thirty)

	// Bitmap flags (OR-able)
	fmt.Println("\n*** IOTA BITMAP FLAGS ***")

	const (
		read    = 1 << iota // 0001 (binary)
		write               // 0010
		execute             // 0100
		isLink              // 1000
	)
	fmt.Println("bit shift (dec):", read, write, execute, isLink, read|execute|isLink)
	fmt.Printf("bit shift (bin): %b %b %b %b %b\n", read, write, execute, isLink, read|execute|isLink)

	// This does not work:
	// fmt.Println("\n*** NO RUNTIME FUNCS FOR CONST INIT ***")

	// const (
	// 	pow2 = math.Pow(2, iota) // error
	// )
}
