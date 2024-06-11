package main

import (
	"errors"
	"fmt"
)

func main() {

	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator) // The caller must be able to handle the err as well checking to see if err is nil or not to perform functionality
	if err != nil {
		fmt.Printf(err.Error()) // format of accessing the error message
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder) //Printf allows for a cleaner usage of variables when writing strings. Each %v
	}

	/*
				Another format of conditional logic: Switch Statement. The break is implied in normal switch statements in Go
				You can also add variables inside of the switch before the brackets to check for conditional values speifically i.e: Is the remainder == 0, or 1,2 and print a statement for each occurence
		switch{
		case err!= nil:
			fmt.Printf(err.error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
		}
	*/
}

func printMe(printValue string) { // Golang takes paramenters and arguments similar to Swift and other languages, type must be included in the parameter
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) { // return type must also be written, you can also return multiple items by writing the types that will return in paranthesis
	var err error // if a func can return an error such as dividing by zero, error handling is handled such as adding an error return value
	// A check must occur that would prompt the error and then return values that make sense to the function
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err // Lastly,the return value in which the function succeeds must return the err variable as it is equal to nil in this case

}
