package main

import "fmt"

// Helloer is an interface.
// An interface type defines a behavior.
// It consists of one or more function signatures.
// Any data type that implements all of these functions
// implements the interface.
type Helloer interface {
	Hello(string)
}

// Greeting implements Hello(string)
type Greeting string

func (g Greeting) Hello(name string) {
	fmt.Println(g+",", name)
}

// Invitation also implements Hello(string)
type Invitation struct {
	event string
}

func (in *Invitation) Hello(name string) {
	fmt.Printf("Welcome to my %s, %s! Come in!\n", in.event, name)
}

func main() {
	// h is an interface. Anything that implements Helloer
	// can be assigned to this variable.
	var h Helloer

	// Make h a Greeting
	h = Greeting("Hello")
	h.Hello("Gopher")

	// Make h an Invitation
	h = &Invitation{event: "birthday party"}
	h.Hello("Kitty")

	// h is still of type Helloer and does not know anything about Invitation's event field.

	// fmt.Println(h.event) // not possible
}
