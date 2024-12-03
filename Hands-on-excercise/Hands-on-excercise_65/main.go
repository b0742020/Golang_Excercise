package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(subtract(10, 20))
	fmt.Println(multiply(10, 20))
	fmt.Println(divide(10, 20))
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}
