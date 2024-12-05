package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop , the ucket gets filled."
	q := "Persistently ,patiently ,you are bond to succeed."
	r := "The meaning of life is ......"
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
