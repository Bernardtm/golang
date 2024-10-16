package main

import (
	"fmt"
	"testing"
)

func Test_multiple_variables(t *testing.T) {
	// Note: If you use the type keyword, it is only possible to declare one type of variable per line.
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// If the type keyword is not specified, you can declare different types of variables in the same line:
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

  // Go Variable Declaration in a Block
  var (
    i int
    j int = 1
    k string = "hello"
  )

 fmt.Println(i)
 fmt.Println(j)
 fmt.Println(k)
}
