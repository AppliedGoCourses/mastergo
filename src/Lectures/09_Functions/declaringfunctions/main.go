package main

import (
	"errors"
	"fmt"
)

func simple() {
	fmt.Println("A simple function")
}

func oneParameter(n int) {
	fmt.Println("A function with a parameter:", n)
}

func multipleParameters(n int, m int, s string) {
	fmt.Println("A function with parameters:", n, m, s)
}

func groupedParameters(n, m int, s string) {
	fmt.Println("Parameters can be grouped by type.")
}

func variadicParameters(s ...string) {
	for _, str := range s {
		fmt.Print(str + " ")
	}
	fmt.Println()
}

func singleReturnValue() int {
	return 99
}

func multipleReturnStatements() int {
	if true {
		return 99
	}
	return 0 // You cannot omit this one.
}

func multipleReturnValues() (int, string, error) {
	return 0, "ok", nil
}

func namedReturnValues() (score int, level string, err error) {
	score = 5000
	level = "master of disaster"
	err = errors.New("guru meditation")
	return score, level, err // a naked return is possible but can lead to confusion.
}

func main() {
	simple()
	oneParameter(37)
	multipleParameters(32, 16, "eight")
	groupedParameters(32, 16, "eight")
	variadicParameters("Clear", "is", "better", "than", "clever")
	fmt.Println(singleReturnValue())
	fmt.Println(multipleReturnStatements())
	n, s, err := multipleReturnValues()
	fmt.Println(n, s, err)
	fmt.Println(namedReturnValues())
}
