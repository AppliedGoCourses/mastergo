// This text is from bufio.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func buf() {

	// Open a file. os.File implements io.Reader.
	f, err := os.Open("bufio.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Create a buffered reader.
	r := bufio.NewReader(f)

	// Now we can use methods that are not available
	// for os.File directly. (Let alone for io.Reader.)
	s, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}
