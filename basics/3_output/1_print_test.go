package basics

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	// Tip: \n creates new lines.
	fmt.Print(j, "\n")
	// printing multiple variables.
	fmt.Print(i, "\n", j)
	// If we want to add a space between string arguments, we need to use " "
	fmt.Print(i, " ", j)
	// Print() inserts a space between the arguments if neither are strings
	fmt.Print(i, j)
}
