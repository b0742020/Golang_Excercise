package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := [5]int{1, 2, 3, 4, 5}
	for k, v := range xi {
		fmt.Println("Index:", k, "Value:", v)
	}
}
