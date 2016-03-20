// Declare a named type called counter with a base type of int. Declare a variable
// named c of type counter set to its zero value. Display the value of c.
//
// Declare a variable named c2 of type counter set to the value of 10. Display the value
// of c2.
//
// Declare a variable named i of the base type int. Attempt to assign the value
// of i to c. Does the compiler allow the assignment?
package main

// Add imports.
import "fmt"

// Declare the named type called counter with a base
// type of int.
type counter int

// main is the entry point for the application.
func main() {
	// Declare a variable named c of type counter
	// set to its zero value.
	var c counter

	// Display the value of c.
	fmt.Println(c)

	// Declare a variable named c2 of type counter
	// set to the value of 10.
	c2 := counter(10)

	// Display the value of c2.
	fmt.Println(c2)

	// Declare a variable named i of the base type int.
	var i int

	// Assign the value of i to the variable named c.
	i = c

	// Does it compile?
	// Nope! Although counter is a type with a base type of int, it is still a different type than int itself. Therefore, i which is of type int, cannot hold the value c, which is of type counter
}
