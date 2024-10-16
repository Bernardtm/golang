package main

import (
	"fmt"
	"testing"
)

// Information can be passed to functions as a parameter. Parameters act as variables inside the function.
// Parameters and their types are specified after the function name, inside the parentheses.
// You can add as many parameters as you want, just separate them with a comma.

// Note: When a parameter is passed to the function, it is called an argument.
// So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.
func familyNameWithSingleParameter(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func TestFunctionWithParameters(*testing.T) {
	familyNameWithSingleParameter("Liam")
	familyNameWithSingleParameter("Jenny")
	familyNameWithSingleParameter("Anja")
	familyNameWithMultipleParameters("Liam", 3)
	familyNameWithMultipleParameters("Jenny", 14)
	familyNameWithMultipleParameters("Anja", 30)
}

func familyNameWithMultipleParameters(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}
