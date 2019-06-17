package main

import (
	"fmt"
	"sort"

	"Lectures/19_interfaces/temperr"
)

func main() {
	fmt.Println("\n*** bufio ***\n")
	buf()

	fmt.Println("\n*** error ***\n")
	numFailures = 3 // Simulate two subsequent failures
	for {
		err := failTemporarily()
		if err == nil {
			break
		}
		fmt.Println("Failed - retrying:", err)
	}
	fmt.Println("Success")

	fmt.Println("\n*** sort ***\n")

	list := List{"really really long", "short", "quite long", "longer"}

	// List implements sort.Interface, so sorting is a snap.
	sort.Sort(list)

	fmt.Printf("%#v\n", list)

}

// to simulate a series of failures, we use this global counter
var numFailures int

func failTemporarily() error {
	numFailures--
	if numFailures > 0 {
		return temperr.New("Temp error")
	}
	return nil
}
