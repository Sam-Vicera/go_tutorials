package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32 // Array initialization. Arrays are of fixed length, of the Same Type, are Indexable, and are contiguous in Memory
	// default values are those of the type being used, in this case
	intArr[1] = 123

	/*
		A few different ways to initialize an array:
		var intArr [3]int32 = [3]int32{1,2,3}
			intArr := [3]int32{1,2,3}
			intArr := [...]int32{1,2,3} --> Let Go infer how many elements there are in an array by using three dots and then listing how many elements are in the array.
	*/

	fmt.Println(intArr[0]) // --> How to access a specific index in an array
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1]) // --> printing out memory allocaiton, an entryway into pointers that will be explored later on
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6} // --> Creation of a slice, a type of wrapper around an Array that adds more funcitonality that an Array doesnt have.

	fmt.Printf("\n The length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7) // adding a value to the slice. Due to the underlying array only having a capacity of 3, equal to the length due to it not being specified, a new array is made when appending

	fmt.Printf("\n The length is %v with capacity %v \n", len(intSlice), cap(intSlice)) // Due to needing an increase in capacity, the initial capacity is doubled when making the new array.

	fmt.Println(intSlice) // while there more capacity in this new slice, the indexes cannot be accessed until a value is received in those new spots

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // A spread operator can be used in such a way to append multiple elements in one go
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 9) // Another way of making a slice in which the first number provided is a required length while the second number is an optional capacity
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8) // initalizing maps
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45} // Creating a map with values pre inserted
	fmt.Println(myMap2["Adam"])                            // If an attempt is made to access an element that does not exist, the return types default value is returned, in this case 0

	var age, ok = (myMap2["Sam"]) // maps in Go have an optional second value of type Bool which allows to check if a key exists otherwise another piece of logic is performed.
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
