package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Name string `json:"name" sql:"NAME"`
	id   int64  // internal
}

func main() {
	m := Mineral{Name: "Pyrite"}

	fmt.Println("\nm:", m)

	// fmt.Println("\n\n***** Struct *****")
	// inspect(m)

	fmt.Println("\n\n***** Pointer *****")
	inspect(&m)

	fmt.Println("\nm:", m)
}

func inspect(intf interface{}) {
	examineType(reflect.TypeOf(intf))
	examineAndSetValue(reflect.ValueOf(intf))
}

func examineType(t reflect.Type) {
	fmt.Println("\n*** Examining the type ***")

	fmt.Println("Type of t:", t)
	fmt.Println("Name of t:", t.Name())
	fmt.Println("Kind of t:", t.Kind())

	isPtr := t.Kind() == reflect.Ptr
	fmt.Println("Is t a pointer?", isPtr)

	if isPtr {
		fmt.Println("Changing t to t.Elem()")
		t = t.Elem()
		fmt.Println("Type of t:", t)
		fmt.Println("Kind of t:", t.Kind())
		fmt.Println("Is t now a struct?", t.Kind() == reflect.Struct)
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

func examineAndSetValue(v reflect.Value) {
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

	name := v.Field(0)
	fmt.Println("Value of field 'Name':", name)

	fmt.Println("Is field Name settable?", name.CanSet())
	fmt.Println("Is internal field id settable?", v.Field(1).CanSet())

	if !name.CanSet() {
		return
	}

	name.SetString("Fool's Gold")
	fmt.Println("New value of field 'Name':", name)

}
