// Created by: Mariam Kasim
// Created on: March 2023
//
// This program gives a random positive or negative number
package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
var randomNum int
var positiveOrNegative string
var positive = "positive"
var negative = "negative"

// input
	fmt.Println("This program gives a random positive or negative number")
	fmt.Println()
	fmt.Println("Would you like to generate a random positive or negative number? (positive/negative):")
	fmt.Scanln(&positiveOrNegative)

	// process
	if positiveOrNegative == positive {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 6
		randomNum = rand.Intn(max-min+1) + min
	} else if positiveOrNegative == negative {
		rand.Seed(time.Now().UnixNano())
		min := -6
		max := -1
		randomNum = rand.Intn(max-min+1) + min
	} else {
		fmt.Println("Invalid input")
	}

	// output
	fmt.Println("The random number is:", randomNum)
	fmt.Println("\nDone.")
}
