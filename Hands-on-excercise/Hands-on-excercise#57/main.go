package main

func main() {
	x := sum([]int{1, 2, 4, 3, 5, 9, 7, 9, 5, 3, 2, 1})
	println(x)
}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}
