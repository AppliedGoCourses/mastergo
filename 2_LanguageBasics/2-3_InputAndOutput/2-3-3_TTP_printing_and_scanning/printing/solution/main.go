package main

import "fmt"

func main() {
	// Print RGB values...
	r, g, b := 124, 87, 3

	// ...as #7c5703  (specifying hex format, fixed width, and leading zeroes)
	fmt.Printf("#%02X%02X%02X\n", r, g, b)

	// ...as rgb(124, 87, 3)
	fmt.Printf("rgb(%d, %d, %d)\n", r, g, b)

	// ...as rgb(124, 087, 003) (specifying fixed width and leading zeroes)
	fmt.Printf("rgb(%03d, %03d, %03d)\n", r, g, b)

	// ...as rgb(48%, 34%, 1%) (specifying a literal percent sign)
	fmt.Printf("rgb(%d%%, %d%%, %d%%)\n", r*100/255, g*100/255, b*100/255)

	// Print the variable type of r
	fmt.Printf("r: %T\n", r)
}
