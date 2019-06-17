package main

import "fmt"

func trace(name string) func() {
	fmt.Println("Entering", name)
	return func() {
		fmt.Println("Leaving", name)
	}
}
func f() {
	defer trace("f")()
	fmt.Println("Doing something")
}
func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
