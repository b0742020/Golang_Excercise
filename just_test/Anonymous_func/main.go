package main

import "fmt"

func main() {
	foo()
	// Anonymous function
	func() {
		fmt.Println("This is anonymous function")
	}()

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))

}

func foo() {
	fmt.Println("This is foo fucntion")
}
