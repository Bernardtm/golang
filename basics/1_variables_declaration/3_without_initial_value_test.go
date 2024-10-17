package basics

import (
	"fmt"
	"testing"
)

// Declaring (Creating) Variables
func TestWithoutInitialValue(t *testing.T) {

	// Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println(a) // ""
	fmt.Println(b) // 0
	fmt.Println(c) // false

	// Value Assignment After Declaration
	var student1 string
	student1 = "John"
	fmt.Println(student1)

}
