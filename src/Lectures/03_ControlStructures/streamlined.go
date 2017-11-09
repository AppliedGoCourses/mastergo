package main

// The streamlined version of nesting.go.
// Note the increase in readability.

func main() {
	var condition1, condition2, condition3 bool
	//...

	if condition1 {
		//...
		return
	}
	if condition2 {
		//...
		return
	}
	if condition3 {
		//...
		return
	}
	// Code of last else branch here.
}
