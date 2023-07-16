package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
