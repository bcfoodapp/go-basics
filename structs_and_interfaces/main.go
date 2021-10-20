package main

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

// This is a "constructor" because the name starts with "new"
// There is no formal definition of a constructor in Go
func NewFoodTruck(foodTruckName string) FoodTruck {
	// Struct initialization
	return FoodTruck{
		name:      foodTruckName,
		openToday: true,
	}
}

// Interfaces for composition relationships
type Vendor interface {
	SetName(name string)
}

func main() {
	// Zero initialization (creates a nil pointer in this case)
	var w Vendor

	// Initializes FoodTruck, takes its pointer address, and stores it as a Vendor value
	// Why does it say "Type does not implement 'Vendor'"?
	w = &FoodTruck{
		name:      "",
		openToday: true,
	}
	w.SetName("El Jefe")
}

/*
https://gobyexample.com/structs
https://gobyexample.com/methods
https://tour.golang.org/basics/3 (On exported names)
https://gobyexample.com/interfaces
*/
