package main

import "fmt"

func changeElement(s []int) {
	s[0] = -1
}

func changeSlice(s []int) {
	s = []int{4, 5, 6}
	fmt.Println("s inside changeSlice:", s)
}}

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println("s1:", s1)
	changeElement(s1)
	fmt.Println("s1 after calling changeElement:", s1)
	changeSlice(s1)
	fmt.Println("s1 after calling ChangeSlice:", s1)
}}
