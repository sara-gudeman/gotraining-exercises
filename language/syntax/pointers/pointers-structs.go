// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
func updateEmail(u *user, update string /* add pointer parameter, add value parameter */) {
	// Use the pointer to change the value that the
	// pointer points to.
	u.email = update
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	sara := user{
		name:  "Sara",
		email: "hello@gmail.com",
		age:   24,
	}

	// Display the value of the variable.
	fmt.Println("email: ", sara.email)

	// Share the variable with the function you declared below.
	updateEmail(&sara, "goodbye@gmail.com")

	// Display the value of the variable.
	fmt.Println("updated email: ", sara.email)
}
