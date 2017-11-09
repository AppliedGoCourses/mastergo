package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Weight int `unit:"g"`
	Name   string
	Opaque bool
}

func examineTypes(m Mineral) {
	t := reflect.TypeOf(m)

	// Reflect upon the type of m:

	fmt.Println("Type of m:", t)
	fmt.Println("Kind of m:", t.Kind())
	fmt.Println("Is m a struct?", t.Kind() == reflect.Struct)

	// Depending on the actual type, reflect.Type provides
	// methods for inspecting that type further.
	// In case of a struct, we can examine its fields:

	field, ok := t.FieldByName("Weight")
	if !ok {
		fmt.Println("Field 'Weight' does not exist")
	}
	fmt.Println("Field name:", field.Name)
	fmt.Println("Field type:", field.Type)

	// field is a reflect.StructField and has a field named "Tag"
	// that represents the field tag. The Get() method
	tag := field.Tag
	fmt.Println("Field tag:", tag)
	unit := tag.Get("unit")
	fmt.Println("Field tag 'unit' contains:", unit)

	fmt.Println()
}

func examineValues(m Mineral) {
	v := reflect.ValueOf(m)

	fmt.Println("Value of m:", v)
	fmt.Println("Type of m:", v.Type())
	fmt.Println("Kind of m:", v.Kind())

	field := v.FieldByName("Weight")
	fmt.Println("Value of field 'Weight':", field)

	// To change a reflection value, the value must
	// a) be addressable, and
	// b) be exported.
	// Addressable means that changing the value would change the
	// original value. The field "Weight" is not settable:
	fmt.Println("Can we set the field 'Weight'?", field.CanSet())

	// This is because we created a copy of the original value when
	// calling reflect.ValueOf(m).
	// Let's get the reflect value by reference instead:
	v = reflect.ValueOf(&m)

	// v is now of type *Mineral:
	fmt.Println("Type of &m:", v.Type())

	// We cannot get the field from the pointer, and (*v) does
	// not work either, as it triggers an error:
	// "invalid indirect of v (type reflect.Value)"
	//
	// Instead, method Elem() does the pointer indirection:
	elem := v.Elem()
	fmt.Println("Type of v.Elem():", elem.Type())

	// Now we can fetch the field and change its value:
	field = elem.FieldByName("Weight")
	fmt.Println("Can we set the field 'Weight' now?", field.CanSet())
	field.SetInt(4)
	fmt.Printf("Changed m.Weight to %v\n", m.Weight)
}

func main() {

	t := Mineral{
		Weight: 3,
		Name:   "a test structure",
		Opaque: true,
	}

	examineTypes(t)
	examineValues(t)
}
