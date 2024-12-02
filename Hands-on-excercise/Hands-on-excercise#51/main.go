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
	xi["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range xi {
		fmt.Println(k)
		for i, value := range v {
			fmt.Println(i, value)
		}
	}
	fmt.Println("------------")
	delete(xi, "no_dr")
	fmt.Println("------------")
	for k, v := range xi {
		fmt.Println(k)
		for i, value := range v {
			fmt.Println(i, value)
		}
	}
}
