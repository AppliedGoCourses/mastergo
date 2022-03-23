package main

import (
	"math"
	"time"
)

type CelestialBody struct {
	Name           string
	Mass           int64 // in 10^21 kg
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

var (
	mercury = Planet{
		CelestialBody: CelestialBody{
			Name:           "Mercury",
			Mass:           330, // * 10^21 kg
			Diameter:       4879,
			RotationPeriod: 1407 * time.Hour,
		},
		Gravity:          3.7,
		HasAtmosphere:    false,
		HasMagneticField: true,
		Satellites:       []string{},
	}

	venus = Planet{
		CelestialBody: CelestialBody{
			Name:           "Venus",
			Mass:           4870, // * 10^21 kg
			Diameter:       12104,
			RotationPeriod: 5833 * time.Hour,
		},
		Gravity:          8.9,
		HasAtmosphere:    true,
		HasMagneticField: false,
		Satellites:       []string{},
	}

	earth = Planet{
		CelestialBody: CelestialBody{
			Name:           "Earth",
			Mass:           5970, // * 10^21 kg
			Diameter:       12756,
			RotationPeriod: 24 * time.Hour,
		},
		Gravity:          9.8,
		HasAtmosphere:    true,
		HasMagneticField: true,
		Satellites:       []string{"Moon"},
	}

	mars = Planet{
		CelestialBody: CelestialBody{
			Name:           "Mars",
			Mass:           642, // * 10^21 kg
			Diameter:       6792,
			RotationPeriod: 24 * time.Hour,
		},
		Gravity:          3.7,
		HasAtmosphere:    true,
		HasMagneticField: false,
		Satellites:       []string{"Phobos", "Deimos"},
	}

	sun = Star{
		CelestialBody: CelestialBody{
			Name:           "Sun",
			Mass:           1988000000,
			Diameter:       1391400,
			RotationPeriod: 587,
		},
		Distance:    0,
		Magnitude:   4.83,
		Discovery:   math.MinInt64, // There is no MinInt, that's why Discovery is an int64
		FirstPlanet: &mercury,
	}
)

func main() {
	// next and previous links have to be instantiated after declaring all variables,
	// in order to avoid circular declarations

	mercury.next = &venus
	venus.next = &earth
	earth.next = &mars

	venus.previous = &mercury
	earth.previous = &venus
	mars.previous = &earth

}
