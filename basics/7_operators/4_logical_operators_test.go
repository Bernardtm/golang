package main

import (
	"fmt"
	"testing"
)

// Logical operators are used to determine the logic between variables or values
func TestAnd(*testing.T) {
	var x = 5
	fmt.Println(x < 5 && x < 10)
}

func TestOr(*testing.T) {
	var x = 5
	fmt.Println(x < 5 || x < 4)
}

func TestNot(*testing.T) {
	var x = 5
	fmt.Println(!(x < 5 && x < 10))
}
