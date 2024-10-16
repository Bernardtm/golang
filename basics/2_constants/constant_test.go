package main

import (
	"fmt"
	"testing"
)

// Constants can be declared both inside and outside of a function
const PI = 3.14 // type inferred

// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
func Test_constant (t *testing.T) {
  const PROGRAM_NAME string = "learning-go"
  fmt.Println(PI)
  fmt.Println(PROGRAM_NAME)

  // Multiple Constants Declaration
  const (
    A int = 1
    B = 3.14
    C = "Hi!"
  )
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
