package main

import (
	"errors"
	"fmt"
	"os"
)

// Some data to write to a file
type Doc struct {
	ID    int
	Title string
	Text  string
}

// ReadFile returns a wrapped error

func WriteDoc(path string, doc Doc) error {
	_, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return fmt.Errorf("Cannot write document '%s' (id %d): %w", doc.Title, doc.ID, err)
	}
	// write data...
	return nil
}

// Let's inspect the error now.

func main() {
	doc := Doc{
		ID:    20,
		Title: "Error Inspection",
		Text:  "In the previous lecture, we learned about wrapping errors...",
	}

	err := WriteDoc("/path/to/no_file", doc)
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
		fmt.Println("pathError info - Op:", pathErr.Op, ", Path:",
			pathErr.Path, "Err:", pathErr.Err)
		fmt.Println("Did we hit a timeout: ", pathErr.Timeout())
	}
}
