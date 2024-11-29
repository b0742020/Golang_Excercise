package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
