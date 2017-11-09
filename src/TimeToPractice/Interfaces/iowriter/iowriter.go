package main

import (
	"fmt"
)

// TerminalWriter writes to a terminal, breaking the lines
// at the given width.
type TerminalWriter struct {
	width int
}

// Write writes slice p to stdout, in chunks of tw.width bytes,
// separated by newline.
// It returns the number of successfully written bytes, and
// any error that occurred.
// If the complete slice is written, Write returns error io.EOF
func (tw *TerminalWriter) Write(p []byte) (n int, err error) {

	// TODO: Implement Write according to the requirements.

	return n, err
}

// NewTerminalWriter creates a new TerminalWriter. width is
// the terminal's width.
func NewTerminalWriter(width int) *TerminalWriter {
	return &TerminalWriter{width: width}
}

func main() {
	s := []byte("This is a long string converted into a byte slice for testing the TerminalWriter.")

	tw := NewTerminalWriter(20)
	n, err := tw.Write(s)
	fmt.Println(n, "bytes written. Error:", err)

}
