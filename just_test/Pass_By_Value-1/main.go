package main

import "fmt"

// value semantics

func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOnePtr(v *int) {
	*v = *v + 1
}

func main() {

	// value semantics

	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	// pointer semantics

	fmt.Println("pointer semantics")
	b := 1
	fmt.Println(b)
	addOnePtr(&b)
	fmt.Println(b)

}
