package main

import "fmt"

//定義結構 name 的型態
type name string

//定義結構 dog
type dog struct {
	first name
}

//定義dog的method
func (d dog) run() {
	fmt.Println("My name is ", d.first, "and I am running.")
}

//定義dog的method
func (d dog) walk() {
	fmt.Println("My name is ", d.first, "and I am walking.")
}

//定義interface 去實現上面兩種方法

type youngin interface {
	walk()
	run()
}

func yongRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Max"}
	yongRun(d1)
	d2 := dog{"Min"}
	yongRun(d2)
	d3 := d1.changeName("Jack")
	yongRun(d3)
}

func (d dog) changeName(n name) dog {
	d.first = n
	return d
}
