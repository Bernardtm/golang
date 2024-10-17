package intermediate

import (
	"fmt"
	"testing"
)

// Go supports methods defined on struct types.
func TestMethods(t *testing.T) {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	// Go automatically handles conversion between values and pointers for method calls.
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
}

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
