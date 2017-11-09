package main

import (
	"fmt"
	"os"
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
func (tw *TerminalWriter) Write(p []byte) (n int, err error) {

	var cnt, offset int
	for more := true; more; {

		// We want to write as many bytes as fit into one terminal line,
		// so we start with the terminal width...
		length := tw.width
		// ...but we can only print as much bytes as there are in p.
		if len(p)-offset < tw.width {
			length = len(p) - offset
			more = false
		}

		cnt, err = os.Stdout.Write(p[offset : offset+length])
		fmt.Println() // Write a newline. Internally, this writes to os.Stdout
		n += cnt
		offset += length
	}
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
