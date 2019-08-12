package main

import (
	"fmt"

	"Lectures/18_TypeAliases/oldlib"
)

func main() {
	var t oldlib.MyType = 10
	oldlib.MyVar = oldlib.MyConst
	fmt.Println(oldlib.MyFunc(t))
}
