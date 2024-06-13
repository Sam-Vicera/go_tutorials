package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

func main() {

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Printf("\nThe price of chicken at %v is %v", website, chickenPrice)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}
func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		fmt.Printf("\nThe price of tofu at %v is %v", website, tofuPrice)
		if tofuPrice <= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select { // switch statement for channels
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found a deal on chicken at %v", website)

	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found a deal on tofu at %v", website)
	}

}

// channels --> Hold Data, Thread Safe, Listen for Data
