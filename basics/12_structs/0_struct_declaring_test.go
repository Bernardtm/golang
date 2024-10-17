package basics

import (
	"fmt"
	"testing"
)

/*
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

To declare a structure in Go, use the type and struct keywords
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
*/

type person struct {
	name string
	age  int
}

// constructor
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func TestStruct(*testing.T) {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don’t have to give it a name.
	// The value can have an anonymous struct type.
	// This technique is commonly used for table-driven tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
