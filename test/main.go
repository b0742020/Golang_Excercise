package main

import "fmt"

func main() {
	x1 := []int{102, 255, 365, 455, 588}
	for k, v := range x1 {
		fmt.Printf("Index: %d, Value: %d\n", k, v)
	}
	/*for i := 0; i < 10; i++ {
		fmt.Println("--")
		for j := 0; j < 10; j++ {
			fmt.Printf("Outer loop: %d, Inner loop: %d\n", i, j)
		}
	}*/
}
