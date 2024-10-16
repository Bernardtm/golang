package main

import (
	"fmt"
	"testing"
)

func Test_array_initialization(t *testing.T) {
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1) // [0 0 0 0 0]
	fmt.Println(arr2) // [1 2 0 0 0]
	fmt.Println(arr3) // [1 2 3 4 5]
	// Tip: The default value for int is 0, and the default value for string is "".
}

func Test_array_specific_initialization(t *testing.T) {
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1) // [0 10 40 0 0]
}

func Test_array_length(t *testing.T) {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
