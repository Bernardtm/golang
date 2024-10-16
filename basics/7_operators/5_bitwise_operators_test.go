package main

import (
	"fmt"
	"testing"
)

// Bitwise operators are used on (binary) numbers
func TestBitwiseAnd(*testing.T) {
	var x = 9
	var y = 8
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("x & y is %b\n", x&y)
}

func TestBitwiseOr(*testing.T) {
	var x = 9
	var y = 8
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("x | y is %b\n", x|y)
}

func TestBitwiseXor(*testing.T) {
	var x = 9
	var y = 8
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("x ^ y is %b\n", x^y)
}

func TestZeroFillLeftShift(*testing.T) {
	var x = 9
	fmt.Printf("x = %b\n", x)
	fmt.Printf("x << 2 is %b\n", x<<2)
}

func TestSignedRightShift(*testing.T) {
	var x = 9
	fmt.Printf("x = %b\n", x)
	fmt.Printf("x >> 2 is %b\n", x>>2)
}
