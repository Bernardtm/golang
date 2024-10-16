package main

import (
	"fmt"
	"testing"
)

/*
  The integer data type has two categories:

  Signed integers - can store both positive and negative values
  Unsigned integers - can only store non-negative values
*/
// Tip: The default type for integer is int. If you do not specify a type, the type will be int.
func Test_signed_integer(*testing.T) {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
	/*
	  Go has five keywords/types of signed integers:

	  Type	Size	Range
	  int	Depends on platform:
	  32 bits in 32 bit systems and
	  64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
	  -9223372036854775808 to 9223372036854775807 in 64 bit systems
	  int8	8 bits/1 byte	-128 to 127
	  int16	16 bits/2 byte	-32768 to 32767
	  int32	32 bits/4 byte	-2147483648 to 2147483647
	  int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807
	*/
}

func Test_unsigned_integer(*testing.T) {
	var x uint = 500
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
	/*
	  Go has five keywords/types of unsigned integers:

	  Type	Size	Range
	  uint	Depends on platform:
	  32 bits in 32 bit systems and
	  64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
	  0 to 18446744073709551615 in 64 bit systems
	  uint8	8 bits/1 byte	0 to 255
	  uint16	16 bits/2 byte	0 to 65535
	  uint32	32 bits/4 byte	0 to 4294967295
	  uint64	64 bits/8 byte	0 to 18446744073709551615
	*/
}
