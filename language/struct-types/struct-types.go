// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	sara := user{
		name:  "Sara",
		email: "hello@gmail.com",
		age:   24,
	}

	// Display the field values.
	fmt.Println("name", sara.name)
	fmt.Println("email", sara.email)
	fmt.Println("age", sara.age)

	// Declare a variable using an anonymous struct.
	david := struct {
		name  string
		email string
		age   int
	}{
		name:  "David",
		email: "goodbye@gmail.com",
		age:   26,
	}

	// Display the field values.
	fmt.Println("name", david.name)
	fmt.Println("email", david.email)
	fmt.Println("age", david.age)
}
