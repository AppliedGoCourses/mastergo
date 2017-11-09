// +build freebsd netbsd openbsd

package main

import "fmt"

func HelloFromOS() {
	fmt.Println("Hello from a BSD system!")
}
