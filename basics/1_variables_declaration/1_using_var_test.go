package main

import (
	"fmt"
	"testing"
)

// Declaring (Creating) Variables - With the var keyword
func TestUsingVar(t *testing.T) {
  // Can be used inside and outside of functions
  // Variable declaration and value assignment can be done separately
  var variablename int = 1
  var student = "Jane" // type is inferred

  // Note: You always have to specify either type or value (or both).
  fmt.Println(student)
  fmt.Println(variablename)

}
