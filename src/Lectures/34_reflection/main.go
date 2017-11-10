package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Name string `json:"name"`
	id   int64  // internal
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

	field, ok := t.FieldByName("Name")
	if !ok {
		fmt.Println("Field not found")
		return
	}
	fmt.Println("Field name:", field.Name)
	fmt.Println("Field type:", field.Type)

	tag := field.Tag
	fmt.Println("Field tag:", tag)
	unit := tag.Get("json")
	fmt.Println("Field tag 'json' contains:", unit)

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

	if v.Kind() != reflect.Struct {
		return
	}

	field := v.FieldByName("Name")
	fmt.Println("Value of field 'Name':", field)
}

func modifyValue(v reflect.Value) {
	fmt.Println("\n*** Modifying the value ***")

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("Expected Struct, got", v.Kind())
		return
	}

	field := v.FieldByName("Name")
	fmt.Println("Old value of field 'Name':", field)
	if !field.CanSet() {
		fmt.Println("Field 'Name' is not settable.")
		return
	}

	field.SetString("Fool's Gold")
	fmt.Println("New value of field 'Name':", field)
}

func handleSomeInterfaceParameter(intf interface{}) {
	examineType(reflect.TypeOf(intf))
	examineValue(reflect.ValueOf(intf))
	modifyValue(reflect.ValueOf(intf))
}

func main() {
	m := Mineral{Name: "Pyrite"}

	fmt.Println("\nm:", m)

	fmt.Println("\n\n***** Struct *****")
	handleSomeInterfaceParameter(m)

	fmt.Println("\n\n***** Pointer *****")
	handleSomeInterfaceParameter(&m)

	fmt.Println("\nm:", m)
}
