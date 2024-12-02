package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "I'm 008"}
	z := [][]string{x, y}
	for k, v := range z {
		fmt.Println(k, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
	fmt.Println(z)
}
