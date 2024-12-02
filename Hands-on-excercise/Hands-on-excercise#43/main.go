package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range xi {
		fmt.Printf("Value: %v\tType: %T\n", v, v)
	}
}
