package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorialLoop(4))
}

func factorial(n int) int {
	fmt.Println("Which number is n right now", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
