package intermediate

import "fmt"

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.

func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0

  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {

  sum(1, 2)
  sum(1, 2, 3)

  // Variadic functions can be called in the usual way with individual arguments.
  // If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
