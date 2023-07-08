package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Println(a[len(a)-1])
	fmt.Printf("len(a): %v\n", len(a))

	fmt.Printf("\nArray a\n")
	for i, v := range a {
		fmt.Printf("i: %d -> v: %d\n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}

	fmt.Printf("\nArray q\n")
	for i, v := range q {
		fmt.Printf("i: %d -> v: %d\n", i, v)
	}

	fmt.Printf("\nArray r\n")
	for i, v := range r {
		fmt.Printf("i: %d -> v: %d\n", i, v)
	}

	s := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T -> %v\n", s, s)

	s = [10]int{0, 1, 2, 3, 4}
	fmt.Printf("%T -> %v\n", s, s)

	symbol := [...]string{USD: "Dollar", EUR: "Euro", GBP: "Libra", RMB: "Yene"}
	fmt.Printf("%T -> %v\n", symbol, symbol)

	fmt.Println(GBP, symbol[GBP])

	R := [...]int{99: -1}
	fmt.Printf("%T -> %v\n", R, R)

	A := [2]int{1, 2}
	B := [...]int{1, 2}
	C := [2]int{1, 3}
	fmt.Println(A == B, A == C, B == C)
	// D := [3]int{1, 2}
	// fmt.Println(A == D) // compile error: cannot compare [2]int == [3]int
}
