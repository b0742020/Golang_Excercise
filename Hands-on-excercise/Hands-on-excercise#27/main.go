package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println("Random number generated is:", x)
	fmt.Println("Random number generated is:", y)
	if x < 4 && y < 4 {
		fmt.Println("Both numbers are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both numbers are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("One number is greater than 4 and less than 6")
	} else if y != 5 {
		fmt.Println("y is not equal to 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}

}
