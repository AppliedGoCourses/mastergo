package main

import (
	"fmt"
	"time"
)

type CelestialBody struct {
	Name           string
	Mass           int64
	Diameter       int64
	Gravity        float64
	RotationPeriod time.Duration
}

type Planet struct {
	HeavenlyBody     CelestialBody
	HasAtmosphere    bool
	HasMagneticField bool
	Satellites       []string
	next, previous   *Planet
}

func main() {

	fmt.Print("\n\n*** Chained field access ***\n\n")

	var p Planet

	// This can become quite verbose, esp. if more nesting
	// levels are involved
	p.HeavenlyBody.Name = "Venus"
	p.HeavenlyBody.Mass = 4.87e15
	p.HeavenlyBody.Diameter = 12104

	fmt.Println("p.HeavenlyBody.Name:", p.HeavenlyBody.Name)

	// Solution: Embed the struct as an anonymous field

	fmt.Print("\n\n*** Embedded struct ***\n\n")

	type Planet2 struct {
		CelestialBody    // Anonymous field: No name, only a type
		HasAtmosphere    bool
		HasMagneticField bool
		Satellites       []string
		next, previous   *Planet
	}

	var p2 Planet2
	p2.Name = "Venus" // shorthand for p2.CelestialBody.Name

	fmt.Println("p2.Name:", p2.Name)

	fmt.Print("\n\n*** A not so anonymous field ***\n\n")

	// Note the use of the type name "CelestialBody" here.
	// The full access path is still available for anonymous fields.
	p2.CelestialBody.Gravity = 8.9

	fmt.Println("p2.CelestialBody.Gravity is", p2.CelestialBody.Gravity)

	fmt.Print("\n\n*** The shortest path wins ***\n\n")

	type Planet3 struct {
		Name             string
		CelestialBody    // Anonymous field: No name, only a type
		HasAtmosphere    bool
		HasMagneticField bool
		Satellites       []string
		next, previous   *Planet
	}

	var p3 Planet3

	// Now the shorthand p3.Name for p3.CelestialBody.Name
	// is not available anymore.
	p3.Name = "Venus"
	fmt.Println("p3.Name:", p3.Name)
	fmt.Println("p3.CelestialBody.Name:", p3.CelestialBody.Name)

	fmt.Print("\n\n*** No special syntax for struct literals ***\n\n")

	// This does not work.
	// p := Planet2{
	// 	Name: "Mercury",  // Error: unknown field 'Name' in struct literal of type Planet2
	// }

	p2 = Planet2{
		CelestialBody: CelestialBody{
			Name: "Mercury",
		},
	}

	fmt.Print("\n\n*** Anonymous pointers ***\n\n")

	type s1 struct {
		name string
	}

	type s2 struct {
		*s1 // pointer!
	}

	str2 := s2{
		&s1{
			name: "What's the point(er)?",
		},
	}

	// Works like expected.
	fmt.Println("str2:", str2)
	fmt.Println("*s1:", *str2.s1) // same as *(str2.s1)

	fmt.Print("\n\n*** Anonymous non-struct types ***\n\n")

	// Any named type can be used as an anonymous field.
	type anon struct {
		int     // anonymous int
		float64 // anonymous float64
	}

	a := anon{}
	a.int = 4
	a.float64 = 4.4
	fmt.Printf("%v\n", a.int)
	fmt.Printf("%v\n", a.float64)
	fmt.Printf("%#+v\n", a)

	fmt.Print("\n\n*** Ambiguous field names in embedded structs ***\n\n")

	type t struct {
		name string
	}

	type u struct {
		name string
	}

	type s struct {
		t
		u
	}

	// Now s.Name is ambiguous.
	// Is it the shorthand for s.t.Name or rather for s.u.Name?
	// The compiler cannot decide this, so it errors out.

	st := s{}

	// This does not compile.
	// st.name = "who am I"
	// fmt.Println(st.name)

	st.t.name = "t"
	st.u.name = "u"
	fmt.Println(st)
}
