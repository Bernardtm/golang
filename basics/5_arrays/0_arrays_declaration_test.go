package main

import (
	"fmt"
	"testing"
)

// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
// Note: The length specifies the number of elements to store in the array. In Go, arrays have a fixed length.
// The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).
func TestArraysUsingVarDefinedLength(*testing.T) {
	// var array_name = [length]datatype{values} // here length is defined
	var arr1_defined_length = [3]int{1, 2, 3}
	fmt.Println(arr1_defined_length)
}

func TestArraysUsingVarInferredLength(*testing.T) {
	// var array_name = [...]datatype{values} // here length is inferred
	var arr1_inferred_length = [...]int{1, 2, 3}
	fmt.Println(arr1_inferred_length)
}

func TestArraysUsingWalrusDefinedLength(*testing.T) {
	// array_name := [length]datatype{values} // here length is defined
	arr2_defined_length := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr2_defined_length)
}

func TestArraysUsingWalrusInferredLength(*testing.T) {
	// array_name := [...]datatype{values} // here length is inferred
	arr2_inferred_length := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr2_inferred_length)
}

func TestArraysStrings(*testing.T) {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)
}
