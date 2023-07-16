package main

import (
	"fmt"
	"strings"
)

// squares returns a function that returns
// the next square number each time it called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {

	s := strings.Map(func(r rune) rune { return r + 1 }, "Vinicius")
	fmt.Printf("s: %v\n", s)

	f := squares()
	fmt.Printf("f: %v\n", f)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	x := squares()
	fmt.Printf("x: %v\n", x)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

	addFive := addUp(5)
	fmt.Printf("addFive(1): %v\n", addFive(1))
	addThree := addUp(3)
	fmt.Printf("addThree(7): %v\n", addThree(7))
}

func addUp(add int) func(int) int {
	return func(i int) int {
		return i + add
	}
}
