package main

import (
	"fmt"
	"testing"
)

func Test_slice_change_element(*testing.T) {
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
