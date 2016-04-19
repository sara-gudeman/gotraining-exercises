// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
	ages := make(map[string]int)

	// Intialize some data into the map.
	ages["Gwen"] = 56
	ages["Cary"] = 6
	ages["Sam"] = 34
	ages["Ken"] = 16

	// Display each key/value pair.
	for key, value := range ages {
		fmt.Println("Name:", key, "Age: ", value)
	}
}
