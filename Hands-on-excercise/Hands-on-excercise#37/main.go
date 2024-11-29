package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("Name : %s\tAge : %d\n", k, v)
	}

	fmt.Println("-----------")

	age1 := m["James"]
	fmt.Printf("Age of James is %d\n", age1)
	if v, ok := m["James"]; ok {
		fmt.Println("There is a BOND lookup entry , and Bond's age is", v)
	}

	age2 := m["Q"]
	fmt.Println("That age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no entry for Q", v)
	}
}
