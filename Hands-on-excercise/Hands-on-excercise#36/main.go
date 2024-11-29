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
}
