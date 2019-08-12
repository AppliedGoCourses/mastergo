package main

import (
	"fmt"
	"log"
)

// A simple and flexible way to create a new error is using `fmt.Errorf()`. All formatting options of fmt.Printf() are available. Errorf() returns an `error` value.
func verify(i int) error {
	// As an example, we test the input against an allowed range of values.
	if i < 0 || i > 10 {
		return fmt.Errorf("verify: %d is outside the allowed range (0..10)", i)
	}
	// If no error occurred, return nil.
	// The type "error" is not a pointer type, but rather an interface type,
	// and interfaces also can be nil.
	return nil
}

// Propagate an error value to the caller.
// A good style is to wrap the error into a new message
// to create a simple trace through the failing call chain.
func propagate(i int) error {
	// "if action; check error" style
	if err := verify(i); err != nil {
		return fmt.Errorf("propagate: %s", err)

		// Or try the convenience library "pkg/errors".
		// Install it by calling
		//     go get github.com/pkg/errors
		// Then uncomment the following line:

		// return errors.Wrap(err, "pkgErrors")
	}
	return nil
}

func retry(i int) error {
	// "action - if err" style
	err := propagate(i)

	// Retry with half of the absolute value.
	if err != nil {
		if i < 0 {
			i = -i
		}
		err = propagate(i / 2)
		// If propagate still errors out, return an error.
		if err != nil {
			return fmt.Errorf("retry: %s", err)
		}
	}
	return nil
}

// In rare cases it might be sufficient to log an error and
// then continue.
func onlyLog(i int) {
	if err := retry(i); err != nil {
		log.Println("onlyLog:", err)
	}
}

// log.Fatal, log.Fatalln, and log.Fatalf write a message
// to the log output and terminate the running process.
func logAndExit(i int) {
	err := retry(i)
	if err != nil {
		log.Fatalln("exit:", err)
	}
}

func main() {

	fmt.Println("\n*** CREATING ERRORS ***\n")

	err := verify(12)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n*** HANDLING ERRORS ***\n")

	// propagate wraps the error into a new one and passes this to the caller.
	if err := propagate(-1); err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	// Not using `if err` anymore to keep the code short.
	// All of the functions error out anyway.
	fmt.Println(retry(24))

	fmt.Println()

	// onlyLog returns no error
	onlyLog(36)

	fmt.Println()

	// Note that log.Println prepends date and time to the log message.
	// You can change this by calling log.SetPrefix() and log.SetFlags().
	// The name of the command makes a good prefix, and the flags can be
	// ORed together (like logSetFlags(log.Lshortfile | log.Ltime)).

	log.SetPrefix("errors: ")
	log.SetFlags(log.Lshortfile)
	onlyLog(48)

	fmt.Println()

	// Terminate the program
	logAndExit(60)
}
