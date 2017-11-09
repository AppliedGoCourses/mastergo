package main

import (
	"errors"
	"fmt"
)

// Step 1. This is our generic data queue.

// Queue is a list of data, with FIFO semantics.
type Queue []interface{}

// PutAny adds an element to the queue.
func (c *Queue) PutAny(elem interface{}) {
	*c = append(*c, elem)
}

// GetAny removes an element from the queue.
// If the queue is empty, return an error.
func (c *Queue) GetAny() (interface{}, error) {
	if len(*c) == 0 {
		return nil, errors.New("empty queue")
	}
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem, nil
}

// Step 2. Write a type-safe wrapper for int values.

type IntQueue struct {
	Queue
}

// Through Put(), nothing but an int type can
// enter the list
func (ic *IntQueue) Put(n int) {
	ic.PutAny(n)
}

// Get returns the type-asserted int value.
func (ic *IntQueue) Get() (int, error) {
	n, err := ic.GetAny()
	if err != nil {
		return 0, err
	}
	return n.(int), nil
}

// The calling code does the type assertion when retrieving an element.
func main() {
	ic := IntQueue{}
	ic.Put(7)
	ic.Put(42)

	for i := 0; i < 3; i++ {
		elem, err := ic.Get()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Got: %d (%[1]T)\n", elem)
	}
}
