package main

import "fmt"

func verify(i int) error {
	if i < 0 || i > 10 {
		return fmt.Errorf("verify: %d is outside the allowed range", i)
	}
	return nil
}

func propagate(i int) error {
	if err := verify(i); err != nil {
		return fmt.Errorf("propagate (%d): %w", i, err)
	}
	return nil
}

func main() {
	fmt.Println(propagate(24))
}
