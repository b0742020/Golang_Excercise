package main

import "fmt"

type x struct {
	first     string
	friends   map[string]int
	favDrinks []string
}

func main() {
	p1 := x{
		first: "John",
		friends: map[string]int{
			"Jane": 25,
			"Bob":  30,
		},
		favDrinks: []string{
			"Coffee",
			"Tea",
			"Water",
		},
	}
	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	p2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Jane",
		friends: map[string]int{
			"John": 35,
			"Bob":  40,
		},
		favDrinks: []string{
			"Coffee",
			"Tea",
			"Milk",
		},
	}
	fmt.Println(p2)
}
