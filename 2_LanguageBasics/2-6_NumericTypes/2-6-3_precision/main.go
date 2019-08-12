package main

import "fmt"

func main() {
	var f64 float64 = 642545754
	var f32 float32 = 642545754

	fmt.Printf("%.0f\n", f64)
	fmt.Printf("%.0f\n", f32)
	fmt.Printf("f32 is off by %.0f\n", f64-float64(f32))
}
