package main

import "fmt"

func main() {
	v := 1
	v++
	v--
	v ^= v
	fmt.Printf("v: %v\n", v)

	x, y := 1, 2
	fmt.Printf("x: %v y: %v\n", x, y)
	x, y = y, x
	fmt.Printf("x: %v y: %v\n", x, y)

	fmt.Printf("gcd(56, 14): %v\n", gcd(56, 14))
	fmt.Printf("fib(10): %v\n", fib(10))

	medals := []string{"gold", "silver", "bronze"}
	/*
		medals[0] = "gold"
		medals[1] = "silver"
		medals[2] = "bronze"
	*/
	fmt.Printf("medals: %v\n", medals)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
