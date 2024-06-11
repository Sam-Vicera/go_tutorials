package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767 // Golang is a strictly written language, every var must be used if it's being declared/initialized
	// Different ints, Ex: int16, int32, int64. These are used in order to specifiy the amount of memory you'd like to use
	// uint	exists as an unsigned int which only takes positive numbers allowing for the usage of double capacity
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9 // Go doesnt have a float type so you have to choose between 32 or 64 bit numbers
	fmt.Println(floatNum)             // good to think about what kind of float number you need based off context, for example, storing rgb numbers

	// arithmetic operations in go can only occur if values are converted to a common value. Example below

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2

	fmt.Println(intNum1 / intNum2) // the answer given is that of an integer rounded down. If the remainder is needed, the % operator should be used
	fmt.Println(intNum1 % intNum2) // --> 0.5 is returned

	var myString string = "Hello" + " " + "World" // strings can be created using double quotes or back ticks. Back ticks allow the string to be initialzied on different lines as shown below.
	// string concatenation is similar to other languages
	// \n can be used to create a line break as well.
	var myString2 string = ` This is another way to
	create a string in Go!
	`
	fmt.Println((myString + myString2))

	fmt.Println(len("test")) // returns the length of a string but not by chars but by the amount of bytes it represents from the ASCll table
	// in order to deal with this issue, we can import a package that allows us to get the direct length of a string as below, check above for the utf package

	fmt.Println(utf8.RuneCountInString("γ")) // --> gamma char which returns two using len but one when using RuneCount

	var myRune rune = 'a'
	fmt.Println(myRune) // runes are a different data type in Go but will be looked into in a future project.

	var myBoolean bool = false
	fmt.Println(myBoolean) // bools are either true or false like other languages

	var intNum3 int
	fmt.Println(intNum3) // vars don't have to be initailzed as Go has default values for strictly written vars

	myVar := "text" // var type can be excluded if the var is initialzed with a value immediately. This is infered initialization
	// var keyword can also be excluded with the syntax shown above
	fmt.Println(myVar)

	var1, var2 := 1, 2 // multi initalization is possible like this
	fmt.Println(var1, var2)

	const myConst string = "const value" // constants of Go, need a value when initalized
	fmt.Println(myConst)
}
