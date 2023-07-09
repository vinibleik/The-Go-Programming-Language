package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("runes: %q\n", runes)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a: %v len(a): %v cap(a): %v\n", a, len(a), cap(a))
	s := a[:]
	fmt.Printf("s: %v len(s): %v cap(s): %v\n", s, len(s), cap(s))
	y := appendInt(s, 10)
	fmt.Printf("a: %v len(a): %v cap(a): %v\n", a, len(a), cap(a))
	fmt.Printf("s: %v len(s): %v cap(s): %v\n", s, len(s), cap(s))
	fmt.Printf("y: %v len(y): %v cap(y): %v\n\n", y, len(y), cap(y))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("b: %v len(b): %v cap(b): %v\n", b, len(b), cap(b))
	t := b[1:4]
	fmt.Printf("t: %v len(t): %v cap(t): %v\n", t, len(t), cap(t))
	x := appendInt(t, 10)
	fmt.Printf("b: %v len(b): %v cap(b): %v\n", b, len(b), cap(b))
	fmt.Printf("t: %v len(t): %v cap(t): %v\n", t, len(t), cap(t))
	fmt.Printf("x: %v len(x): %v cap(x): %v\n", x, len(x), cap(x))

	var j, k []int
	for i := 0; i < 10; i++ {
		k = appendInt(j, i)
		fmt.Printf("%d len=%d cap=%d\t%v\n", i, len(k), cap(k), k)
		j = k
	}

	var p []int
	p = appendInts(p, 1)
	p = appendInts(p, 2, 3)
	p = appendInts(p, 4, 5, 6)
	p = appendInts(p, p...) // append the slice p
	fmt.Println(p)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}
