package main

import (
	"fmt"
	"time"
)

// Person holds details of a person.
type Person struct {
	Name        string    // exported
	dateOfBirth time.Time // internal
}

// Age is a package-level function for returning the age of a person.
func Age(p Person) int {
	return int(time.Since(p.dateOfBirth).Hours() / 24 / 365) // ignoring leap days
}

// Age as a method of Person.
func (p Person) Age() int {
	return int(time.Since(p.dateOfBirth).Hours() / 24 / 365)
}

// ChangeName changes the receiver's Name field.
// Here, we need a pointer to the receiver; otherwise, p would just
// be a copy of the original Person struct.
func (p *Person) ChangeName(name string) {
	p.Name = name // remember, p.Name is a shorthand for (*p).Name
}

// ChangeNothing uses a non-pointer receiver, but in this case, the change
// only affects the copy p and not the original receiver.
func (p Person) ChangeNothing(name string) {
	p.Name = name // remember, p.Name is a shorthand for (*p).Name
	fmt.Println("Inside ChangeNothing: p.Name is", p.Name)
}

func main() {
	fmt.Println("\n\n*** Methods ***\n\n")

	p := Person{Name: "Gordon", dateOfBirth: time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)}

	// A method call is similar to accessing a struct field
	name := p.Name // struct field access
	age := p.Age() // method call
	fmt.Printf("Name: %s, age:%d\n", name, age)

	// If you want to update the receiver, use a pointer
	p.ChangeName("Glenda")
	fmt.Println("Name changed through pointer receiver:", p.Name)

	// Calling a method with pointer receiver on a pointer
	(&p).ChangeName("Gopher")
	fmt.Println("(&p).ChangeName:", p.Name)

	// Calling a method with non-pointer receiver on a pointer also works
	(&p).Age()
	fmt.Println("(&p).Age()", p.Age())

	// To summarize all pointer/non-pointer combinations
	fmt.Println("\n\n*** Pointer/non-pointer receiver vs pointer/non-pointer variable ***\n\n")

	pp := &Person{"Glenda", time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)} // pointer to Person

	fmt.Println("NR-NV:", p.Age())  // non-pointer receiver (NR), non-pointer variable (NV)
	fmt.Println("NR-PV:", pp.Age()) // non-pointer receiver, pointer variable (PV)
	p.ChangeName("The Gopher")      // pointer receiver (PR), non-pointer variable
	pp.ChangeName("Gordon")         // pointer receiver, pointer variable
	fmt.Println("PR-NV:", p.Name)
	fmt.Println("PR-PV:", pp.Name)

	fmt.Println("\n\n*** A non-pointer receiver is only a copy of the original value (pass-by-value) ***\n\n")

	p.ChangeNothing("Phoger")
	fmt.Println("ChangeNothing: ", p.Name)

	fmt.Println("\n\n*** Embedding: Call embedded methods through shorthand access ***\n\n")

	type Contact struct {
		Person
		Email string
	}

	c := Contact{p, "gopher@go.lang"}
	fmt.Println("Embedded method Age():", c.Age()) // shorthand for c.Person.Age()

}
