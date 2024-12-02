package main

import "fmt"

func main() {
	x1 := []int{102, 255, 365, 455, 588}
	fmt.Println(x1)
	x1 = append(x1[:2], x1[3:]...)
	fmt.Println(x1)

	am := map[string]int{
		"Gibson": 24,
		"Pee":    20,
	}
	fmt.Println(am)
	fmt.Println("The age of Gibson is", am["Gibson"])

	an := make(map[string]int)
	an["John"] = 25
	an["Mary"] = 22
	fmt.Println(an)

	for k, v := range am {
		fmt.Println(k, v)
	}
	for _, v := range an {
		fmt.Println(v)
	}
	fmt.Println("----------")
	delete(an, "John")
	fmt.Println(an)
	fmt.Println(an["John"])

	fmt.Println("----------")
	v, ok := am["Gibson"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key not found")
	}
}
