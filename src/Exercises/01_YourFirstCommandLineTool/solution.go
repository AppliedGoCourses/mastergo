// This is one solution for exercise 01, "Build your first commandline tool".
package main

// Import all required packages. An editor with a decent Go plugin can
// automatically add import statements for packages used in the code,
// as long as the packages reside in the standard library or are found
// in $GOPATH/src.
import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// A floating-point variable that holds the value to convert.
	var input float64

	// A floating-point value that holds the result.
	var output float64

	// Set up and read the -from and -to flags.
	// The arguments to flag.StringVar() are: A pointer to the target
	// variable, the name, the default value, and a description.
	var from, to string
	flag.StringVar(&from, "from", "", "The unit to convert from")
	flag.StringVar(&to, "to", "", "The unit to convert to")
	flag.Parse()

	// Let's look how the arguments look like...

	//...from the os package's perspective
	fmt.Println("os.Args:", os.Args)

	//...and from the flag package's perspective (after having parsed the flags):
	fmt.Println("flag.Args():", flag.Args())

	// Check the -from and -to flags. If empty, print an error message
	// and exit.
	if from == "" {
		fmt.Println("Please provide a unit to convert from.")
		return
	}
	if to == "" {
		fmt.Println("Please provide a unit to convert to.")
		return
	}

	// Check if an argument is passed. If not, print a
	// message and exit.
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a value to convert.")
		return
	}

	// Scan the input value (remember this should be the only remaining
	// value in flag.Args() after the flags were parsed).
	//
	// fmt.Sscanf is the variant of Scanf that scans from a string.
	// Its arguments are: The string to scan, the format string,
	// and one or more pointer to the variable(s) to fill.
	//
	// flag.Args() is a slice of strings. Use the index operator [n]
	// to fetch the string at position n.
	numScanned, err := fmt.Sscanf(flag.Args()[0], "%f", &input)
	if numScanned != 1 || err != nil {
		fmt.Println("Cannot scan the value to convert. flag.Args() is:", flag.Args())
		return
	}

	// Convert the value based on the from and to units.
	// Here, we use a switch statement to select the formula
	// depending on the from and to units.
	// Let's keep things simple and only convert from:
	//
	// * meters to feet,
	// * feet to meters,
	// * meters to inches, and
	// * inches to meters.
	//
	switch {
	case from == "meters" && to == "feet":
		output = input / 0.3048
	case from == "feet" && to == "meters":
		output = 0.3048 * input
	case from == "meters" && to == "inches":
		output = input / 0.0254
	case from == "inches" && to == "meters":
		output = 0.00254 * input
	}

	// Finally, print the result.
	//

	fmt.Printf("%f %s are %f %s\n", input, from, output, to)
}
