package main

import (
	"fmt"
	"testing"
)

func TestSliceChangeElement(*testing.T) {
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
