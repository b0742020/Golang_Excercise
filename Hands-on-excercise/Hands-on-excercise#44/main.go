package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for k, v := range xi {
		fmt.Println(k, v)
	}
	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])
}
