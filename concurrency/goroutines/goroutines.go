// Create a program that declares two anonymous functions. One that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import "runtime"

// Add imports.
import "fmt"
import "sync"

// init is called prior to main.
func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

// main is the entry point for the application.
func main() {
	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		for i := 100; i >= 0; i-- {
			fmt.Println("Decrement from 100 ", i)
		}

		// Decrements the count of the wait group.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		for i := 0; i <= 100; i++ {
			fmt.Println("Increment to 100", i)
		}

		// Decrements the count of the wait group.
		wg.Done()
	}()

	fmt.Println("Waiting for program to finish")

	// Wait for the goroutines to finish.
	wg.Wait()

	fmt.Println("Program complete!")
}
