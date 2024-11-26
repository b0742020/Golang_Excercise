package main

import "fmt"

var a int = 10

const b string = "Hello World"

func main() {
	fmt.Printf("a is %d\t%T\n", a, a)
	fmt.Printf("b is %s\t%T\n", b, b)
	c := 20
	fmt.Printf("The value of c is %d\n", c)
}
