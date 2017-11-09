package main

import (
	"fmt"
	"strings"
	"time"
)

type Planet struct {
	Name             string
	Mass             int64
	Diameter         int
	Gravity          float64
	RotationPeriod   time.Duration
	HasAtmosphere    bool
	HasMagneticField bool
	Satellites       []string
	next, previous   *Planet
}

func main() {

	fmt.Print("\n*** Declaration ***\n\n")

	// Like any other variable declaration.
	var earth, jupiter Planet

	// All fields are initialized to their zero value.
	fmt.Printf("Zero value of earth: %v\n\n", earth)

	// A struct literal.
	mars := Planet{
		Name:           "Mars",
		Mass:           642e15, // in millon metric tons (1t == 1000kg)
		Diameter:       6792,
		RotationPeriod: 24.7 * 60 * time.Minute,
		HasAtmosphere:  true,
		Satellites:     []string{"Phobos", "Deimos"},
		previous:       &earth,
		next:           &jupiter, // Remember the final comma
	}

	fmt.Printf("The mars struct: %v\n\n", mars)
	fmt.Printf("The mars struct, with field names: %+v\n\n", mars) // Note: +v

	fmt.Print("\n\n*** Access ***\n\n")

	// Dot access: stuct.field
	mars.Gravity = 3.7
	fmt.Println("Dot access - Mars Gravity:", mars.Gravity, "m/s^2")

	// Dot access even works with a pointer to a struct
	var pmars = &mars
	fmt.Println("Dot access through pointer:", pmars.Gravity)
	// same effect: fmt.Println((*pmars).Gravity)

	fmt.Print("\n\n*** Visibility ***\n\n")

	// Field names starting with an uppercase letter denote an exported field.

	fmt.Println("Mars diameter:", mars.Diameter, "km")

	// This would not work if called from outside the package:
	fmt.Println("Next planet:", mars.next)

	fmt.Print("\n\n*** Comparison ***\n\n")

	// This does not work.
	// fmt.Println("Is Mars the same as Jupiter?", mars == jupiter)

	// invalid operation: mars == jupiter (struct containing []string cannot be compared)

	// Let's make a comparable struct.
	type CelestialBody struct {
		Name           string
		Mass           int64
		Diameter       int64
		Gravity        float64
		RotationPeriod time.Duration
	}

	var sun, moon CelestialBody

	sun.Name = "Sun"
	moon.Name = "Moon"

	// sun and moon have different names.
	// This compiles and returns false.
	fmt.Println("Are sun and moon the same?", sun == moon)

	fmt.Print("\n\n*** Passing to and returning from functions ***\n\n")

	// A function that takes a Planet.
	// Scroll down for the function definition.
	fmt.Println("Does Mars have satellites?", hasSatellites(mars))

	// A function that takes and returns a planet.
	// Note the pass-by-value semantics. `mars.Name` is still the old name,
	// because `uppercase()` receives a copy of `mars`
	fmt.Println(uppercase(mars).Name, "is uppercase of", mars.Name) // Note how we can use dot access on a function result

	// A function that takes a pointer to Planet can modify the struct directly
	rename(&mars, "Ἄρεως ἀστἡρ")
	fmt.Println("An ancient Greek name for mars is", mars.Name)

	fmt.Print("\n\n*** Struct literal without field names (not recommended) ***\n\n")

	// Can you tell from memory which value belongs to which field?
	mars = Planet{
		"Mars",
		642e15,
		6792,
		3.7,
		24.7 * 60 * time.Minute,
		true,
		false,
		[]string{"Phobos", "Deimos"},
		&earth,
		&jupiter,
	}

	fmt.Printf("Created without using field names: %+v\n", mars)

}

func hasSatellites(p Planet) bool {
	return len(p.Satellites) > 0
}

func uppercase(p Planet) Planet {
	p.Name = strings.ToUpper(p.Name)
	return p
}

func rename(p *Planet, n string) {
	p.Name = n
}
