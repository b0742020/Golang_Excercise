package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	favIC     []string
}

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		favIC:     []string{"apple", "banana", "orange"},
	}

	p2 := person{
		firstName: "Jane",
		lastName:  "Doe",
		favIC:     []string{"pear", "grape", "watermelon"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.favIC)
	fmt.Println(p2.favIC)
	for k, v := range p1.favIC {
		fmt.Println(k, v)
	}
	for k, v := range p2.favIC {
		fmt.Println(k, v)
	}

}
