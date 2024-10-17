package basics

import (
	"fmt"
	"testing"
)

func TestAssignment(*testing.T) {
	var x = 10
	fmt.Println(x)
}

func TestAdditionAssignment(*testing.T) {
	var x = 10
	x += 5
	fmt.Println(x)
}

func TestSubtractionAssignment(*testing.T) {
	var x = 5
	x -= 3
	fmt.Println(x)
}

func TestMultiplicationAssignment(*testing.T) {
	var x = 5
	x *= 3
	fmt.Println(x)
}

func TestDivisionAssignment(*testing.T) {
	var x = 5
	x /= 3
	fmt.Println(x)
}

func TestModulusAssignment(*testing.T) {
	var x = 5
	x %= 3
	fmt.Println(x)
}

func TestBitwiseAndAssignment(*testing.T) {
	var x = 5
	fmt.Printf("x is %b \n", x)   // 101
	fmt.Printf("3 is %03b \n", 3) // 011
	x &= 3
	fmt.Printf("x now is %03b \n", x) // 001
}

func TestBitwiseOrAssignment(*testing.T) {
	var x = 5
	fmt.Printf("x is %b \n", x)   // 101
	fmt.Printf("3 is %03b \n", 3) // 011
	x |= 3
	fmt.Printf("x now is %03b \n", x) // 111
}

func TestBitwiseXorAssignment(*testing.T) {
	var x = 5
	fmt.Printf("x is %b \n", x)   // 101
	fmt.Printf("3 is %03b \n", 3) // 011
	x ^= 3
	fmt.Printf("x now is %03b \n", x) // 110
}

func TestRightShiftAssignment(*testing.T) {
	var x = 5
	fmt.Printf("x is %b \n", x) // 101
	x >>= 3
	fmt.Printf("x now is %03b \n", x) // 000
}

func TestLeftShiftAssignment(*testing.T) {
	var x = 5
	fmt.Printf("x is %b \n", x) // 101
	x <<= 3
	fmt.Printf("x now is %03b \n", x) // 101000
}
