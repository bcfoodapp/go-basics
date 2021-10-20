package main

import "fmt"

func main() {
	number := 10 // Stack allocated because it is not used outside of main()
	myNewNumber := thisFunctionReturns20(&number)
	fmt.Println(*myNewNumber)

	// myNewNumber is garbage collected
	// What happens if you returned a pointer in C++?
	// https://gobyexample.com/pointers
}

// Pointer types: *int
// Return value written after parameters
func thisFunctionReturns20(number *int) *int {
	// Dereference the pointer
	fmt.Println(*number)
	returnValue := 20
	return &returnValue
	// Return value might be moved to the heap because returnValue is no longer in the stack
}
