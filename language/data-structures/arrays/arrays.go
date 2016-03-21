// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings set to its zero value.
	var strings [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	names := [5]string{"Sam", "Erik", "Jo", "Jeff", "Psy"}

	// Assign the populated array to the array of zero values.
	strings = names

	// Iterate over the first array declared.
	for i, name := range strings {
		// Display the string value and address of each element.
		fmt.Println(name, &strings[i])
	}
}
