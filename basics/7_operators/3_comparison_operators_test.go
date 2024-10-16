package main

import (
	"fmt"
	"testing"
)

// Comparison operators are used to compare two values.
// Note: The return value of a comparison is either true (1) or false (0).
func Test_equal(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x == y)
}

func Test_not_equal(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x != y)
}

func Test_greater_than(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x > y)
}

func Test_less_then(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x < y)
}

func Test_greater_than_equal_to(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x >= y)
}

func Test_less_thenequal_to(*testing.T) {
	var x = 5
	var y = 3
	fmt.Println(x <= y)
}
