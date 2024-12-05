package main

import "fmt"

func intDelta(n *int) {
	*n = *n + 1
}

func sliceDelta(n []int) {
	n[0] = 99
}

func mapDelta(n map[string]int, s string) {
	n[s] = 123
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)
	sliceDelta(b)
	fmt.Println(b)

	c := make(map[string]int)
	c["key"] = 42
	fmt.Println(c)
	mapDelta(c, "key")
	fmt.Println(c)
}
