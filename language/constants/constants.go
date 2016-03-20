// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

// Declare a constant named server of kind string and assign a value.
const server = "sara"

// Declare a constant named port of type integer and assign a value.
const port int = 30

// main is the entry point for the application.
func main() {
	// Display the value of both server and port.
	fmt.Println("server: ", server, "\nport: ", port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	division := 2 / 3.14

	// Display the value of the variable.
	fmt.Println("value: ", division)
}
