package main

import (
	"fmt"
	"testing"
)

// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.
/*
In Go, there are several ways to create a slice:

Using the []datatype{values} format
Create a slice from an array
Using the make() function
*/
func TestSliceUsingArrayLikeSyntax(*testing.T) {
	/*
		  In Go, there are two functions that can be used to return the length and capacity of a slice:

		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/
	// slice_name := []datatype{values}
	myslice1 := []int{} // an empty slice of 0 length and 0 capacity
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"} // a slice of strings of length 4 nd also the capacity of 4
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
