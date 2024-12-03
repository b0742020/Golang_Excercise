package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()
}
