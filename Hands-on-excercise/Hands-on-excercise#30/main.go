package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	x := rand.Intn(5)
	fmt.Println("Random number generated is:", x)
	switch {
	case x < 5:
		fmt.Println("Range between 0 and 4")
	default:
		fmt.Println(" Is an number")
	}
	for i := 1; i < 43; i++ {
		fmt.Println(i)
		x := rand.Intn(5)
		fmt.Println("Random number generated is:", x)
		switch x {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("four")
		}
		fmt.Println("--")
	}
}
