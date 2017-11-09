package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Weight int `unit:"g"`
}

func yesno(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func examineType(t reflect.Type) {
	fmt.Println("\n*** Examining the type ***")

	fmt.Println("Type of t:", t)
	fmt.Println("Kind of t:", t.Kind())
	isPtr := t.Kind() == reflect.Ptr
	fmt.Println("Is t a pointer?", yesno(isPtr))

	if isPtr {
		fmt.Println("Changing t to t.Elem()")
		t = t.Elem()
		fmt.Println("Type of t:", t)
		fmt.Println("Kind of t:", t.Kind())
		fmt.Println("Is t now a struct?", yesno(t.Kind() == reflect.Struct))
	}

	if t.Kind() != reflect.Struct {
		return
	}

	field, ok := t.FieldByName("Weigh")
	if !ok {
		fmt.Println("Field not found")
		return
	}
	fmt.Println("Field name:", field.Name)
	fmt.Println("Field type:", field.Type)

	tag := field.Tag
	fmt.Println("Field tag:", tag)
	unit := tag.Get("unit")
	fmt.Println("Field tag 'unit' contains:", unit)

}

func examineValue(v reflect.Value) {
	fmt.Println("\n*** Examining the value ***")

	if v.Kind() == reflect.Ptr {
		fmt.Println("Changing v to v.Elem()")
		v = v.Elem()
	}

	fmt.Println("Value of v:", v)
	fmt.Println("Type of v:", v.Type())
	fmt.Println("Kind of v:", v.Kind())

	_, ok := v.Interface().(Mineral)
	if ok {
		field := v.FieldByName("Weight")
		fmt.Println("Value of field 'Weight':", field)
	}
}

func handleSomeInterfaceParameter(intf interface{}) {
	examineType(reflect.TypeOf(intf))
	examineValue(reflect.ValueOf(intf))
}

func main() {
	m := Mineral{Weight: 10}

	fmt.Println("\n\n***** Struct *****")
	handleSomeInterfaceParameter(m)

	fmt.Println("\n\n***** Pointer *****")
	handleSomeInterfaceParameter(&m)
}
