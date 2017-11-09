package newlib

import "math/rand"

const MyConst = 299792458

type MyType int

var MyVar MyType

func MyFunc(n MyType) int {
	return int(n)
}

func InitMyVar() {
	MyVar = MyType(rand.Intn(100))
}

func init() {
	InitMyVar()
}
