package main

import (
	"fmt"
	"testing"
)

// The make() function can also be used to create a slice.
func Test_slice_make_defining_capacity(*testing.T) {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))   // 5
	fmt.Printf("capacity = %d\n", cap(myslice1)) // 10
}

func Test_slice_make_omitting_capacity(*testing.T) {
	// Note: If the capacity parameter is not defined, it will be equal to length.
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))   // 5
	fmt.Printf("capacity = %d\n", cap(myslice2)) // 5
}
