package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner --> This syntax is possible however it adds another level to reaching the fields inside that struct such as. myEngine.ownerInfo.name (a little lengthy)
	owner // --> This instead allows us to be able to make .name a field that is abled to be accessed by the gasEngine struct
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 { // method for structure to calculate how many miles a car has left
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 { // method for structure to calculate how many miles a car has left
	return e.mpkwh * e.kwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type engine interface { // allows the usage of any engine type as long as it follows the criteria of having the milesLeft() method that has no parameters and returns a uint8
	milesLeft() uint8
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 12, owner: owner{name: "Sam"}} // An intro to structs, the fields (mpg and gallons) can be omitted and initialized in order
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	var myEngine2 = struct { // This is an example of an annonymous structure. Less reusability and has to be redefined if it needs to be used again (creating a new structure)
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.mpg, myEngine.gallons)

	fmt.Printf("Total miles left in tank %v\n", myEngine.milesLeft()) // similar to instantiating a class in other languages and then calling a method

	canMakeIt(myEngine, 123)
}
