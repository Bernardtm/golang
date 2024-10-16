package main

import (
	"fmt"
	"testing"
)

func Test_change_element(t *testing.T) {
	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)
}
