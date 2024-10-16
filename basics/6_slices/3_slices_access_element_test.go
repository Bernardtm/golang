package main

import (
	"fmt"
	"testing"
)

// You can access a specific slice element by referring to the index number.
func Test_slice_access_element(*testing.T) {
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
