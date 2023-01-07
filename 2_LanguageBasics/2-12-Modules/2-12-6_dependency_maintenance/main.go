package main

import (
	"fmt"
	"strings"

	a "github.com/AppliedGoCourses/A"
	b "github.com/AppliedGoCourses/B"
)

func main() {
	fmt.Println(a.A())
	fmt.Println(b.B())

	if strings.Index(b.B(), "version v") == -1 {
		fmt.Println("Version string must start with a 'v'!")
	}
}
