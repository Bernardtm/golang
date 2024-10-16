package main

import (
	"fmt"
	"testing"
)

func TestSliceFromArray(*testing.T) {
	/*
				var myarray = [length]datatype{values} // An array
		    myslice := myarray[start:end] // A slice made from the array
	*/
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	// In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.
	// The slice can grow to the end of the array. This means that the capacity of the slice is 4.
}
