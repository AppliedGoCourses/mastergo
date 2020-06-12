package diceprinter

import (
	"fmt"

	"github.com/appliedgocourses/dice"
)

func PrintRoll(sides int, string comment) {
	fmt.Printf("%s: %d\n", comment, dice.Roll(sides))
}
