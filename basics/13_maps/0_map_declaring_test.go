package main

import (
	"fmt"
	"testing"
)

/*
Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

Note: The order of the map elements defined in the code is different from the way that they are stored.
The data are stored in a way to have efficient data retrieval from the map.
*/

func TestMapUsingVar(*testing.T) {
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Printf("a\t%v\n", a)
}

func TestMapUsingWalrusOperator(*testing.T) {
	// b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("b\t%v\n", b)
}

func TestMapUsingMake(*testing.T) {
	// var a = make(map[KeyType]ValueType)
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty

	// b := make(map[KeyType]ValueType)
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

// There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.
func TestMapEmpty(*testing.T) {
	// var a map[KeyType]ValueType
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}

/*
Allowed Key Types
The map key can be of any data type for which the equality operator (==) is defined. These include:
Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)

Invalid key types are:
Slices
Maps
Functions

These types are invalid because the equality operator (==) is not defined for them.

Allowed Value Types
The map values can be any type.
*/
