package main

import (
	"fmt"
	"strings"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%sf(%d): %d\n", strings.Repeat(" ", n*4), n, i)
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	go f(0)
	f(1)
	time.Sleep(1 * time.Second)
}
