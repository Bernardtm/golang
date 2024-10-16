package main

import (
	"fmt"
	"testing"
)

func Test_addition(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x + y)
}

func Test_subtraction(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x - y)
}

func Test_multiplication(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x * y)
}

func Test_division(*testing.T) {
	x := 10
	y := 2
	fmt.Println(x / y)
}

func Test_modulus(*testing.T) {
	x := 8
	y := 3
	fmt.Println(x % y)
}

func Test_increment(*testing.T) {
	x := 10
	x++
	fmt.Println(x)
}

func Test_decrement(*testing.T) {
	x := 10
	x--
	fmt.Println(x)
}
