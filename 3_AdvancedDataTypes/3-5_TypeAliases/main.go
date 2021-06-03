package main

import (
	"fmt"

	"github.com/AppliedGoCourses/mastergo/3_AdvancedDataTypes/3-5_TypeAliases/oldlib"
)

func main() {
	var t oldlib.MyType = 10
	oldlib.MyVar = oldlib.MyConst
	fmt.Println(oldlib.MyFunc(t))
}
