package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func init() {
	fmt.Println("Initializing the program")
}

func (p person) speak() {
	fmt.Printf("My name is  %s\n", p.first)
	fmt.Printf("I am %d years old\n", p.age)
}

func main() {
	p1 := person{"John", 25}
	p1.speak()
}
