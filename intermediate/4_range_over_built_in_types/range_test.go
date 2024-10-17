package intermediate

import (
	"fmt"
	"testing"
)

func TestRangeOverBuiltInTypes(*testing.T) {
	nums := []int{2, 3, 4}
	sum := 0
	// range on arrays and slices provides both the index and value for each entry.
	// Below we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	// See Strings and Runes for more details.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
