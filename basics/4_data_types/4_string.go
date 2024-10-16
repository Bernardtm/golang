package main

import (
	"fmt"
	"testing"
)

/*
  The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:
*/
// Tip: The default type for integer is int. If you do not specify a type, the type will be int.
func TestString(*testing.T) {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
