package basics

import (
	"fmt"
	"testing"
)

// Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition
// The developer should be careful with recursion functions as it can be quite easy to slip into writing a function which never terminates,
// or one that uses excess amounts of memory or processor power.
// However, when written correctly recursion can be a very efficient and mathematically-elegant approach to programming.
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func TestRecursion(*testing.T) {
	testcount(1)
}
