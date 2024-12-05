package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Initializing the program")
}

func main() {
	fmt.Println("Starting the program")
	timeFunc(doWork)
	fmt.Println("Program ended")
}

func doWork() {
	for i := 0; i < 1_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
