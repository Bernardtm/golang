package main

import (
	"fmt"
	"testing"
)

func TestPrintln(t *testing.T) {
  var i,j string = "Hello","World"
  // The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end
  fmt.Println(i,j)
}
