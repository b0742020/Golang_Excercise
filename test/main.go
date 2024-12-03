package main

import "fmt"

type ListNode struct{
	Val  int
	Next *ListNode
}

type person struct {
	fist string
}

type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
}
func (head *ListNode,val int)*LisListNode{
	dummy := &ListNode{}
	dummy.Next = head
	cur :=
}
func (p person) speak() {
	fmt.Println("I am", p.fist)
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent", sa.fist)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			fist: "John",
		},
		ltk: true,
	}
	p1 := person{
		fist: "John",
	}
	p2 := person{
		fist: "Jenny",
	}
	// sa1.speak()
	// p1.speak()
	// p2.speak()

	saySomething(sa1)
	saySomething(p1)
	saySomething(p2)
}
