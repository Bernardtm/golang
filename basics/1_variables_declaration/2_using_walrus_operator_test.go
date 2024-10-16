package main

import (
	"fmt"
	"testing"
)

// Declaring (Creating) Variables - With the := sign
func TestWalrusOperator(t *testing.T) {

  // Note: It is not possible to declare a variable using ":=" without assigning a value to it.
  // 	Can only be used inside functions
  // Variable declaration and value assignment cannot be done separately (must be done in the same line)
  variablename2 := 1 // type is inferred
  fmt.Println(variablename2)
}
