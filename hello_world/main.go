package main // Executables must have the package name "main"
// Libraries can have whatever package name they want

// To run: go run ./hello_world

// Imported libraries
import "fmt"

// Entry point
func main() {
	// Variable initialization
	// Use snakeCase
	// No semicolons!
	helloString := "Hi"
	// Variable assignment
	helloString = "Hello world!"
	helloWorld(helloString)

	arraysAndSlices()
}

// Function and parameter(s)
func helloWorld(helloString string) {
	if len(helloString) > 0 {
		fmt.Println(helloString)
	}

	/*
		Go by Example references:
		https://gobyexample.com/hello-world
		https://gobyexample.com/variables
		https://gobyexample.com/functions
	*/
}

func arraysAndSlices() {
	// Stack allocated array (constant size!)
	stringArray := [3]string{"a", "b", "c"}

	fmt.Println()
	for i := 0; i < len(stringArray); i++ {
		fmt.Println(i, stringArray[i])
	}
	fmt.Println()

	// Heap allocated slice (dynamic size!)
	// Slice = list
	stringSlice := make([]string, 0)
	stringSlice = append(stringSlice, "a")
	stringSlice = append(stringSlice, "b")
	stringSlice = append(stringSlice, "c")

	// For range
	for index, item := range stringArray {
		fmt.Println(index, item)
	}
	fmt.Println()

	/*
		https://gobyexample.com/arrays
		https://gobyexample.com/slices
		https://gobyexample.com/range
	*/
}
