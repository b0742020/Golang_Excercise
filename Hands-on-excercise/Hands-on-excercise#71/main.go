package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, n int) string {
	x := f(n)
	return fmt.Sprintf("The square of %d is %d", n, x)
}
func main() {
	fmt.Println(printSquare(square, 5))
}
