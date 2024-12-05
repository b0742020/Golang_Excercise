package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNewNode(val int) *ListNode {
	return &ListNode{Val: val,
		Next: nil}
}

func main() {
	head := createNewNode(1)
	fmt.Printf("%v\n", head.Val)
}
