package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Name   string
	Weight int `unit:"g"`
	Opaque bool
}

func examineTypes(intf interface{}) {
	t := reflect.TypeOf(intf)

	// Reflect upon the type of m:

	fmt.Println("Type of intf:", t)
	fmt.Println("Kind of intf:", t.Kind())
	fmt.Println("Is intf a struct?", t.Kind() == reflect.Struct)

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

func examineValues(intf interface{}) {
	v := reflect.ValueOf(intf)

	fmt.Println("Value of intf:", v)
	fmt.Println("Type of intf:", v.Type())
	fmt.Println("Kind of intf:", v.Kind())

	field := v.FieldByName("Weight")
	fmt.Println("Value of field 'Weight':", field)

	// To change a reflection value, the value must
	// a) be addressable, and
	// b) be exported.
	// Addressable means that changing the value would change the
	// original value. The field "Weight" is not settable:
	fmt.Println("Is intf addressable?", v.CanAddr())
	fmt.Println("Can we modify intf?", v.CanSet())
	fmt.Println("Can we modify the field 'Weight'?", field.CanSet())

	// This is because we created a copy of the original value when
	// calling reflect.ValueOf(intf).
	// Let's get the reflect value by reference instead:
	v = reflect.ValueOf(&intf)

	// v is now of type *interface{}:
	fmt.Println("Type of &intf:", v.Type())

	// We cannot get the field from the pointer (through v.FieldByName(...),
	// and (*v).FieldByName() does not work either, as it triggers an error:
	// "invalid indirect of v (type reflect.Value)"
	//
	// Method Elem() does the required pointer indirection:
	elem := v.Elem()
	fmt.Println("Type of v.Elem():", elem.Type())
	fmt.Println("Is v.Elem() addressable?", elem.CanAddr())
	fmt.Println("Can we modify v.Elem()?", elem.CanSet())

	// Now we can fetch the field and change its value:
	field = elem.FieldByName("Weight")
	fmt.Println("Can we modify the field 'Weight' now?", field.CanSet())
	field.SetInt(4)
	fmt.Printf("Changed intf.Weight to %v\n", intf.(Mineral).Weight)
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
