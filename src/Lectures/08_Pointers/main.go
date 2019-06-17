package main

import "fmt"

func main() {

	// Pointer basics

	var a int = 1
	b := 2 // short declaration; b will be an int, too

	// The address operator
	fmt.Println("\n*** THE ADDRESS OPERATOR ***")

	fmt.Println("a's address:", &a)
	fmt.Println("b's address:", &b)

	// Creating a pointer to a variable
	fmt.Println("\n*** CREATING A POINTER ***")

	var p *int
	fmt.Println("The zero value of p is:", p)

	p = &a
	fmt.Println("p's value is a's address:", p)
	fmt.Println("*p yields a's value:", *p)

	p = &b
	fmt.Println("Now p contains b's address:", p, "and *p is:", *p)

	// The & operator works only with addressable entities.
	// Literals and constants are examples of non-addressable entities.
	//
	// p = &123  // error: "cannot take the address of 123"
	//
	// const c int = 123
	// p = &c  // error: "cannot take the address of c"

	// Modifying a variable through a pointer
	fmt.Println("\n*** MODIFYING A VARIABLE ***")

	*p = 3
	fmt.Println("b was changed to:", b)

	// Pointer indirection triggers a panic for nil pointers.
	// fmt.Println("\n*** NIL POINTER ACCESS ***")
	//
	// p = nil
	// fmt.Println("A nil pointer:", p, *p)  // "panic: runtime error: invalid memory address or nil pointer dereference"

	// Two pointers pointing to the same variable
	fmt.Println("\n*** TWO POINTERS TO A ***")

	p1 := &a
	p2 := &a

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("*p2 is:", *p2)
	*p1 = 4
	fmt.Println("*p2 is:", *p2)

	// Pointer comparisons
	fmt.Println("\n*** POINTER COMPARISON ***")

	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("*p1 == *p2:", *p1 == *p2)

	p = nil
	fmt.Println("p == nil:", p == nil)

	p = &a
	if p != nil {
		*p = 5
	}

	// The new function
	fmt.Println("\n*** THE NEW FUNCTION ***")

	p = new(int)
	fmt.Println("p points to an unnamed int:", p, *p)
	*p = 6
	fmt.Println("The unnamed int has been changed to:", *p)

	// No pointer arithmetic
	// fmt.Println("\n*** NO POINTER ARITHMETIC ***)
	//
	// p = (&a) + 64  // error: "invalid operation: &a + 64 (mismatched types *int and int)"

}
