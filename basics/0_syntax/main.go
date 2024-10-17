// In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.
// Note: In Go, any executable code belongs to the main package.
package basics

// import ("fmt") lets us import files included in the fmt package.
import (
	"fmt"
)

// func main() {} is a function. Any code inside its curly brackets {} will be executed.
func main() { // The left curly bracket { cannot come at the start of a line.
	// fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".
	fmt.Println("Hello World!")
	// In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".
	// Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).
}
