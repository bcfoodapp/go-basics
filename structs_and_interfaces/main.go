package main

import "fmt"

// This is a struct
type FoodTruck struct {
	// This is a field
	name      string
	openToday bool
}

// This is a method
func (p *FoodTruck) IsOpen() bool {
	return p.openToday
}

// What is an alternative way to pass the FoodTruck struct to our method?

// This is a "constructor" because the name starts with "new"
// There is no formal definition of a constructor in Go
func NewFoodTruck(foodTruckName string) *FoodTruck {
	// Struct initialization
	return &FoodTruck{
		name:      foodTruckName,
		openToday: true,
	}
}

// Interfaces for composition relationships (Not object inheritance)
// An interface is a set of required methods
type Vendor interface {
	GetName() string
}

func main() {
	// Zero initialization (creates a nil pointer in this case)
	var vendor Vendor

	// Check the compiler error: go run ./structs_and_interfaces
	vendor = NewFoodTruck("El Jefe")
	fmt.Println(vendor.GetName())

	// Run debugger after the error is fixed to see how the variables change
}

/*
https://gobyexample.com/structs
https://gobyexample.com/methods
https://tour.golang.org/basics/3 (On exported names)
https://gobyexample.com/interfaces
*/
