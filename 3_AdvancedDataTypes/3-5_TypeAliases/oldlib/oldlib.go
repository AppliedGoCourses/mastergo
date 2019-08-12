package oldlib

import (
	"Lectures/18_TypeAliases/newlib"
)

const MyConst = newlib.MyConst

type MyType = newlib.MyType

var MyVar newlib.MyType

func MyFunc(n MyType) int {
	return newlib.MyFunc(n)
}

func init() {
	newlib.InitMyVar()
}
