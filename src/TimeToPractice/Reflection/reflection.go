package main

import (
	"fmt"
	"math"
	"time"

	"TimeToPractice/Reflection/pretty"
)

type CelestialBody struct {
	Name           string
	Mass           int64 `unit:"10^21 kg"`
	Diameter       int64 // in km
	RotationPeriod time.Duration
}

type Planet struct {
	CelestialBody
	Gravity          float64
	HasAtmosphere    bool
	HasMagneticField bool
	Satellites       []string
	next, previous   *Planet
}

type Star struct {
	CelestialBody
	Distance    float64 // in light years
	Magnitude   float64
	Discovery   int64 // year
	FirstPlanet *Planet
}

func main() {
	var mercury, venus, earth, mars Planet
	var sun Star

	mercury = Planet{
		CelestialBody: CelestialBody{
			Name:           "Mercury",
			Mass:           330,
			Diameter:       4879,
			RotationPeriod: 1407 * time.Hour,
		},
		Gravity:          3.7,
		HasAtmosphere:    false,
		HasMagneticField: true,
		Satellites:       []string{},
		next:             &venus,
	}

	venus = Planet{
		CelestialBody: CelestialBody{
			Name:           "Venus",
			Mass:           4870,
			Diameter:       12104,
			RotationPeriod: 5833 * time.Hour,
		},
		Gravity:          8.9,
		HasAtmosphere:    true,
		HasMagneticField: false,
		Satellites:       []string{},
		next:             &earth,
		previous:         &mercury,
	}

	earth = Planet{
		CelestialBody: CelestialBody{
			Name:           "Earth",
			Mass:           5970,
			Diameter:       12756,
			RotationPeriod: 24 * time.Hour,
		},
		Gravity:          9.8,
		HasAtmosphere:    true,
		HasMagneticField: true,
		Satellites:       []string{"Moon"},
		next:             &mars,
		previous:         &venus,
	}

	mars = Planet{
		CelestialBody: CelestialBody{
			Name:           "Mars",
			Mass:           642,
			Diameter:       6792,
			RotationPeriod: 24*time.Hour + 37*time.Minute,
		},
		Gravity:          3.7,
		HasAtmosphere:    true,
		HasMagneticField: false,
		Satellites:       []string{"Phobos", "Deimos"},
		next:             nil,
		previous:         &earth,
	}

	sun = Star{
		CelestialBody: CelestialBody{
			Name:           "Sun",
			Mass:           1988000000,
			Diameter:       1391400,
			RotationPeriod: 587 * time.Hour,
		},
		Distance:    0,
		Magnitude:   4.83,
		Discovery:   math.MinInt64,
		FirstPlanet: &mercury,
	}

	// Not so good: Flat printout of structs.
	fmt.Printf("%v\n", sun)
	fmt.Printf("%v\n", mercury)
	fmt.Printf("%v\n", venus)
	fmt.Printf("%v\n", earth)
	fmt.Printf("%v\n", mars)

	// TODO: Write a library function "pretty.Print(i interface{})"
	// that prints structs in a pretty way.
	// Optional:
	// * Follow pointers
	// * Pretty-print maps
	//
	// Call it as:
	pretty.Print(sun)

}
