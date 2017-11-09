package main

import "fmt"

func main() {

	// String values cannot be changed.
	s1 := "I am immutable"

	// s2 refers to the same string as s1
	s2 := s1

	// s3 is a slice of s1.
	s3 := s1[5:14]

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	fmt.Println()

	// Now let's change s1 and see what happens to s2 and s3
	s1 = "I am new, and " + s1 + ", too"

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	fmt.Println()

	// This is discussed in more detail in the lecture about slices.

	// Create a slice of bytes
	var m1 []byte

	// Convert a string to []byte
	m1 = []byte("I am mutable.")

	// Assign m1 to m2
	m2 := m1

	fmt.Println("m1:", string(m1)) // Conversion of []byte to string
	fmt.Println("m2:", string(m2))

	fmt.Println()

	// Now let's change m1 and see what happens to m2
	m1[2] = 'A' // Single character within single quotes is a byte
	m1[3] = 'M'

	fmt.Println("m1:", string(m1))
	fmt.Println("m2:", string(m2))

}
