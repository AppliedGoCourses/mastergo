package main

import (
	"errors"
	"fmt"
	"os"
)

// ReadFile returns a wrapped error

func ReadFile(path string) (string, error) {
	_, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %w", path, err)
	}
	// more file reading...
	return "file contents", nil
}

// Let's inspect the error now.

func main() {
	s, err := ReadFile("no_such_file")
	if err != nil {

		// The error can be unwrapped once
		fmt.Println("UNWRAP")
		fmt.Println("Top-level error:", err)
		unwrapped := errors.Unwrap(err)
		fmt.Println("Unwrapped error:", unwrapped)

		// Inspect the error: What IS it?

		fmt.Println("\nIS")
		fmt.Println("err is an os.ErrNotExist error:", errors.Is(err, os.ErrNotExist))

		fmt.Println("\nAS")
		var pathErr *os.PathError
		// Treat the error AS an os.PathError error
		fmt.Println("err is an os.PathError:", errors.As(err, &pathErr))
		fmt.Println("pathError info: ", pathErr.Op, "~",
			pathErr.Path, "~", pathErr.Err)
		fmt.Println("Did we hit a timeout: ", pathErr.Timeout())

		return
	}
	fmt.Println(s)
}
