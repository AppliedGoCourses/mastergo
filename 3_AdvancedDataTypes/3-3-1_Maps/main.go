package main

import "fmt"

func main() {

	fmt.Println("\n*** Create a map ***")

	// Variable declaration
	var m map[string]bool
	fmt.Println("m is nil:", m == nil)

	// Instantiating a map using make()
	m = make(map[string]bool)
	fmt.Printf("New empty map: %#v\n", m)
	fmt.Println("m is nil:", m == nil)

	// Pre-allocate space for 100 elements
	m = make(map[string]bool, 100)

	// A map literal
	moons := map[string]int{
		"Earth":   1,
		"Mars":    2,
		"Jupiter": 67,
		"Saturn":  62,
	}
	fmt.Println(moons)
	fmt.Println("moons has", len(moons), "elements.")

	fmt.Println("\n*** Using a non-comparable type as key ***")

	// Example: a slice as a key

	k1 := []string{"Einstein", "Albert"}
	k2 := []string{"Feynman", "Richard"}

	// Make the key type a string instead...
	sliceMap := make(map[string]int)

	// ...and write a function to turn the key value into a string.
	// Let's create an anonymous function and assign it
	// to a variable named "key", to have the function definition and
	// its usage in one place. A standard function definition outside
	// main would do as well.

	key := func(s []string) string {
		return fmt.Sprintf("%v", s)
	}

	sliceMap[key(k1)] = 1879
	sliceMap[key(k2)] = 1918

	fmt.Println("\n*** Insert ***")

	moons["Uranus"] = 27
	moons["Neptune"] = 14
	moons["Pluto"] = 5
	fmt.Println(moons)

	fmt.Println("\n*** Retrieve ***")

	jupiterMoons := moons["Jupiter"]
	fmt.Println("Jupiter has", jupiterMoons, "moons.")

	// Retrieving an element that does not exist returns the zero value.
	mercuryMoons := moons["Mercury"]
	fmt.Println("Mercury has", mercuryMoons, "moons.") // And that's true.

	fmt.Println("\n*** Test for existence ***")

	venusMoons, ok := moons["Venus"]
	if !ok {
		fmt.Println("Data for Venus does not exist. Venus probably has", venusMoons, "moons.")
	}

	_, ok = moons["Venus"]
	fmt.Println("Venus has been added:", ok)

	// pVenusMoons := & moons["Venus"])  // error: cannot take the address of moons["Venus"]

	fmt.Println("\n*** Update ***")

	moons["Jupiter"] = moons["Jupiter"] - 1
	fmt.Println("Jupiter now has only", moons["Jupiter"], "moons!")

	moons["Jupiter"]++
	fmt.Println("False alarm. Jupiter still has", moons["Jupiter"], "moons.")

	fmt.Println("\n*** Delete an element ***")

	delete(moons, "Pluto") // oops, Pluto is not a planet (anymore)

	fmt.Println("\n*** Pass-By-Value vs. internal pointers ***")

	// The add function is defined after main().
	// Even though moons is passed by value (as everything in Go),
	// the copy inside add() still refers to the same elements via pointers.
	fmt.Println("Before calling add():", moons)
	add(moons, "Mercury", 0)
	add(moons, "Venus", 0)
	fmt.Println("After calling add():", moons)

	// newMap is also defined after main().
	fmt.Println("Before calling newMap():", moons)
	newMap(moons)
	fmt.Println("After calling newMap():", moons)

	fmt.Println("\n*** Range loop ***")

	// The output may or may not be ordered by the sequence
	// of inserting the elements.
	// The order in which a range loop reads the elements of a map
	// can even vary each time the loop runs.

	// Run this code multiple times and watch the output of this loop.
	for k, v := range moons {
		fmt.Println("Planet", k, "has", v, "moons.")
	}
}

// add inserts a key-value pair (k, v) into map m.
// The map is changed at the caller's end, too.
func add(m map[string]int, k string, v int) {
	m[k] = v
}

// newMap changes m to refer to a new map.
// m is only a copy of the map at the callers's end, which remains unaffected by the change to m.
func newMap(m map[string]int) {
	m = make(map[string]int)
}
