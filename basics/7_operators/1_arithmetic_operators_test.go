package main

import (
	"fmt"
	"testing"
)

func TestAddition(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x + y)
}

func TestSubtraction(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x - y)
}

func TestMultiplication(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x * y)
}

func TestDivision(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x / y)
}

func TestModulus(*testing.T) {
	x := 8
	y := 3
	fmt.Println(x % y)
}

func TestIncrement(*testing.T) {
	x := 10
	x++
	fmt.Println(x)
}

func TestDecrement(*testing.T) {
	x := 10
	x--
	fmt.Println(x)
}
