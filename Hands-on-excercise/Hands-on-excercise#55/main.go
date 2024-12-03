package main

import "fmt"

type engine struct {
	erletric bool
}
type vehicle struct {
	engine
	make  string
	model string
	door  string
	color string
}

func main() {
	car1 := vehicle{
		engine: engine{
			erletric: true,
		},
		make:  "Toyota",
		model: "Camry",
		door:  "4",
		color: "Red",
	}
	fmt.Println(car1)
	car2 := vehicle{
		engine: engine{
			erletric: true,
		},
		make:  "Toyota",
		model: "Supra",
		door:  "2",
		color: "Red",
	}
	fmt.Println(car2)
}
