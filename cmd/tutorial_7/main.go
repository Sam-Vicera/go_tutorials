package main

import (
	"fmt"
)

func main() {
	// var p *int32 // --> How to initialize a pointer
	var p *int32 = new(int32) // --> How to initialize a pointer, if you try to reference a nil pointer you'll end up with a runtime error
	var i int32

	fmt.Printf("The value of p points to is: %v", *p) // --> dereferencing the pointer, the value that appears is the default value of the type, in this case 0 because the pointer is initialized with int32
	fmt.Printf("\nThe value of i is: %v", i)
	*p = 10 // star syntax has two purposes, the first shown at the top and the second here where it gives a value to what the pointer is pointing at.

	p = &i // Another way of setting a value to a pointer, in this case, setting the value of the location of i
	*p = 1 // if the value of p is changed while having been connected to i, i's value will also change

	fmt.Printf("\nThe value of p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

// slices use pointers under the hood
