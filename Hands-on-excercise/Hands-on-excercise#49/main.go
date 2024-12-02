package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	xi := make(map[string][]string)
	xi["bond_james"] = []string{"shaken", "not stirred", "martinis", "fast cars"}
	xi["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	xi["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	fmt.Printf("%#v", xi)
	for k, v := range xi {
		fmt.Println(k)
		for i, value := range v {
			fmt.Println(i, value)
		}
	}
}
