package basics

import (
	"fmt"
	"testing"
)

// The for loop loops through a block of code a specified number of times.
// The for loop is the only loop available in Go.
func TestForLoop(*testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.
// for index, value := range array|slice|map {
func TestForLoopUsingRange(*testing.T) {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	// Tip: To only show the value or the index, you can omit the other output using an underscore (_).
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
}

// The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
func TestForLoopUsingContinue(*testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// The break statement is used to break/terminate the loop execution.
func TestForLoopUsingBreak(*testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
