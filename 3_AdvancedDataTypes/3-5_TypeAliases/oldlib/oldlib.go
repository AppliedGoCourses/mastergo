package oldlib

import (
	"github.com/AppliedGoCourses/mastergo/3_AdvancedDataTypes/3-5_TypeAliases/newlib"
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
