package main

func main() {

	var motion string

	// Code that sets motion...

	// Simple cases
	switch motion {
	case "walk":
		// Do something
	case "run":
		// Do something else
	case "stop":
		// Do something different
	default:
		// Do nothing
	}

	var velocity int

	// Code that sets velocity...

	// comparison cases
	switch {
	case velocity == 0 || velocity == 1:
		//...
	case velocity >= 10:
		//...
	case f(velocity) >= f(100):
		//...
	}

	// The fallthrough directive
	switch {
	case velocity == 0 || velocity == 1:
		//...
	case velocity >= 10:
		//...
		fallthrough
	case f(velocity) >= f(100):
		//...
	}

	// Multiple case expressions
	switch motion {
	case "walk", "run":
		// Do something
	case "stop":
		// Do something different
	default:
		// Do nothing
	}

}

func f(v int) int {
	return 2 * v
}
