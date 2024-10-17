package intermediate

import (
	"fmt"
	"testing"
)

func TestAnonymousFunction(*testing.T) {

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function!")
	}()
}
