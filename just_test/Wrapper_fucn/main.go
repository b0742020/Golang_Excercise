package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}
