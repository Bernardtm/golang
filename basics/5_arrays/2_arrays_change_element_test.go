package basics

import (
	"fmt"
	"testing"
)

func TestChangeElement(t *testing.T) {
	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)
}
