package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	x := 20
	y := 10
	for {
		fmt.Println("--")
		fmt.Println(x)
		x--
		if x == 0 {
			break
		}
	}
	for y > 0 {
		fmt.Println(y)
		y--
	}
}
