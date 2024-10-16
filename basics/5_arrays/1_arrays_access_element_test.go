package main

import (
	"fmt"
	"testing"
)

func TestAccessElement(t *testing.T) {
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}
