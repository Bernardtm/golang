package main

import (
	"fmt"
	"testing"
)

// If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function
func myFunction(x int, y int) int {
	return x + y
}

// Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name)
func myFunctionWithNamedReturnValueAndNakedReturn(x int, y int) (result int) {
	result = x + y
	return
}

// Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name)
func myFunctionWithNamedReturnValue(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunctionWithMultipleReturnValues(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func TestFunctionReturns(*testing.T) {
	fmt.Println(myFunction(1, 2))
	total := myFunctionWithNamedReturnValue(1, 2)
	fmt.Println(total)
	a, b := myFunctionWithMultipleReturnValues(5, "Hello")
	// If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
	_, c := myFunctionWithMultipleReturnValues(5, "Hello")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
