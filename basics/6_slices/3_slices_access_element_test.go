package main

import (
	"fmt"
	"testing"
)

// You can access a specific slice element by referring to the index number.
func TestSliceAccessElement(*testing.T) {
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
