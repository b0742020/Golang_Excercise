package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for k, v := range xi {
		fmt.Printf("Index: %d, Value: %d\n", k, v)
	}
}
