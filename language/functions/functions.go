// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func createUser(name string, email string, age int) (*user, error) /* (pointer return arg, error return arg) */ {
	// Create a value of type user and return the proper values.
	u := user{
		name:  name,
		email: email,
		age:   age,
	}

	return &u, nil
}

// main is the entry point for the application.
func main() {
	// Use the function to create a value of type user. Check
	// the error being returned.
	u, err := createUser("Sara", "hello@gmail.com", 24)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value that the pointer points to.
	fmt.Println("value: ", *u)

	// Call the function again and just check the error.
	_, err2 := createUser("David", "goodbye@gmail.com", 26)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
