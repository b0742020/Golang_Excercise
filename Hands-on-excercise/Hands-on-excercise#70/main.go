package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	f := whatever()
	fmt.Println(f())
}

func whatever() func() int {
	return func() int {
		return 42
	}
}
