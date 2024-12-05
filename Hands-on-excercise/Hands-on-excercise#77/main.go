package main

import "fmt"

type name string

type person struct {
	first name
}

func main() {
	p := person{first: "John"}
	fmt.Println(p)
	p = changePersonName(p, "Jane")
	fmt.Println(p)
	changePersonNamePtr(&p, "Bob")
	fmt.Println(p)
}

func changePersonName(p person, newFirst name) person {
	p.first = newFirst
	return p
}
func changePersonNamePtr(p *person, newFirst name) {
	p.first = newFirst

}
