package main

import (
	"fmt"
	"testing"
)

// Comparison operators are used to compare two values.
// Note: The return value of a comparison is either true (1) or false (0).
func TestEqual(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x == y)
}

func TestNotEqual(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x != y)
}

func TestGreaterThan(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x > y)
}

func TestLessThan(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x < y)
}

func TestGreaterThanEqualTo(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x >= y)
}

func TestLessThanEqualTo(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x <= y)
}
