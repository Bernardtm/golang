package basics

import (
	"fmt"
	"testing"
)

/*
  The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

  The float data type has two keywords:

  Type	Size	Range
  float32	32 bits	-3.4e+38 to 3.4e+38.
  float64	64 bits	-1.7e+308 to +1.7e+308.

  Tip: The default type for float is float64. If you do not specify a type, the type will be float64.
*/
// Tip: The default type for integer is int. If you do not specify a type, the type will be int.
func TestFloat32(*testing.T) {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}

func TestFloat64(*testing.T) {
	var x float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x)
}
