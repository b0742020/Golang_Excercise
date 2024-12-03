package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(nums ...int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}
func bar(nums []int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}
