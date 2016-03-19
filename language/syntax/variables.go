// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/1xUWjHMB3I

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.

// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).

// main is the entry point for the application.
package main

import "fmt"

func main() {
	// Declare variables that are set to their zero value.
	var name string
	var age int
	var over21 bool

	// Display the value of those variables.
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(over21)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	greeting := "hello"
	date := 19
	isSaturday := true

	// Display the value of those variables.
	fmt.Println(greeting)
	fmt.Println(date)
	fmt.Println(isSaturday)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Println(pi)
}
