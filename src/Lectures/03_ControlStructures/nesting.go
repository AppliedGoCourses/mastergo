package main

func main() {
	var condition1, condition2, condition3 bool
	// some code here...

	// Nesting if/else blocks can decrease readability.
	// See streamlined.go for how to restructure the
	// code to make it cleaner and more readable.

	if condition1 {
		//...
	} else {
		if condition2 {
			//...
		} else {
			if condition3 {
				//...
			} else {
				//...
			}
			// No further code here...
		}
		// ...nor here...
	}
	// ...nor here.
}
