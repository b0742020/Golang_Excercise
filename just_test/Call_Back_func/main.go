package main

import "fmt"

func main() {

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)
	x := doMath(2, 3, add)
	y := doMath(5, 2, sub)

	println(x)
	println(y)

}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {

	return a - b
}
