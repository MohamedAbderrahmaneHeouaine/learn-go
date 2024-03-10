package main

import "fmt"

// to declare a lot of global variable
var (
	var1 = ""
	var2 = ""
)

// to declare a lot of global constant
const (
	First = 1
	Last  = 2
)

func variables() {
	//there are four ways to declare a variable
	// these variable are local in the main function
	var name = "Foo" // infer string
	var firstName string = "Foo"
	var lastName string = "full"
	fullName := firstName + lastName
	//constant var
	const version = 1
	// we can use the first three-way to declare constant

	fmt.Println(name, firstName, fullName)
}
