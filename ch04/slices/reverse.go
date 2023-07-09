package main

import (
	"fmt"
)

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("a: %v\nlen(a): %v\tcap(a): %v\n\n", a, len(a), cap(a))
	reverse(a[:])
	fmt.Printf("a: %v\nlen(a): %v\tcap(a): %v\n\n", a, len(a), cap(a))

	// Rotate s left by two positions.
	s := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s[:])
	fmt.Println(s)

	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("b: %v\n", b)
	rotateLeft(b[:], -1)
	fmt.Printf("b: %v\n", b)
	rotateRight(b[:], -1)
	fmt.Printf("b: %v\n", b)

	x := []int{1, 2, 3, 4, 5, 99: 100}

	fmt.Printf("x: %v\nType(x): %[1]T\tlen(x): %v\tcap(x): %v\n", x, len(x), cap(x))
}

func rotateLeft(s []int, n int) {
	if n < 0 {
		rotateRight(s, -n)
		return
	}

	n %= cap(s)

	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateRight(s []int, n int) {
	if n < 0 {
		rotateLeft(s, -n)
		return
	}

	n %= cap(s)

	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	// fmt.Printf("s: %v\nlen(s): %v\tcap(s): %v\n\n", s, len(s), cap(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
