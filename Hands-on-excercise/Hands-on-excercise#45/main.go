package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	x1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 = append(x1, 52)
	fmt.Println(x1)
	x1 = append(x1, 53, 54, 55)
	fmt.Println(x1)
	y := []int{56, 57, 58, 59, 60}
	x1 = append(x1, y...)
	fmt.Println(x1)
}
