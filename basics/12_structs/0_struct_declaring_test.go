package main

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

type PersonExample struct {
	name   string
	age    int
	job    string
	salary int
}
