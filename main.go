package main

import "fmt"

// to declare a lot of global variable
var (
	var1 = ""
	var2 = ""
)

func main() {
	//there are four ways to declare a variable
	// these variable are local in the main function
	var name = "Foo"
	var firstName string = "Foo"
	var lastName string = "full"
	fullName := firstName + lastName
	fmt.Println(name, firstName, fullName)
}
