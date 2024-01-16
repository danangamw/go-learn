package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age

	// dereference pointer
	fmt.Println("Age: ", *agePointer)

	// using function with pointer
	getAdultYears(agePointer)

	fmt.Println("Adult Years: ", age)
}

// using pointer in function
func getAdultYears(age *int) int {
	// return *age - 18

	// this will change age value
	*age = *age - 18
	return *age
}
