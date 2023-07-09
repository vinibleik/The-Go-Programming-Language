package main

import (
	"unicode"
)

func main() {
	// a := []int{0, 1}
	// fmt.Printf("a: %v\nlen(a): %v\tcap(a): %v\n\n", a, len(a), cap(a))
	// reverse(&a)
	// fmt.Printf("a: %v\nlen(a): %v\tcap(a): %v\n\n", a, len(a), cap(a))

	// b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Printf("b: %v\n", b)
	// rotateLeft(b[:], 9)
	// fmt.Printf("b: %v\n", b)
	// rotateRight(b[:], 3)
	// fmt.Printf("b: %v\n", b)

	// strings := []string{"vini", "go", "go", "test", "baggio", "baggio", "programming"}
	// fmt.Printf("strings: %v\n", strings)
	// fmt.Printf("removeDuplicates(strings): %v\n", removeDuplicates(strings))
	// fmt.Printf("strings: %v\n", strings)

	// strings := []string{"vini", "go", "go", "test", "baggio", "baggio", "programming"}
	// strings = []string{"vini", "vini", "baggio"}
	// fmt.Printf("strings: %v\n", strings)
	// fmt.Printf("removeDuplicates(strings): %v\n", removeDuplicates2(strings))
	// fmt.Printf("strings: %v\n", strings)

	// bytes := []byte("  \t\tde   baggio \n\n         vinicius   \n\n\n\t")
	// fmt.Printf("bytes: %q\n", bytes)
	// fmt.Printf("bytes: %q\n", squashSpaces(bytes))
}

func reverse(arr *[]int) {
	for i, j := len(*arr)-1, 0; i > j; i, j = i-1, j+1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func rotateLeft(arr []int, n int) {
	if n < 0 {
		rotateRight(arr, -n)
		return
	}

	n %= len(arr)

	tail := make([]int, n)
	copy(tail, arr[:n])
	copy(arr[:len(arr)-n], arr[n:])
	copy(arr[len(arr)-n:], tail)
}

func rotateRight(arr []int, n int) {
	if n < 0 {
		rotateLeft(arr, -n)
		return
	}

	n %= len(arr)

	tail := make([]int, n)
	copy(tail, arr[len(arr)-n:])
	copy(arr[n:], arr[:len(arr)-n])
	copy(arr[:n], tail)
}

func removeDuplicates(s []string) []string {
	if len(s) <= 1 {
		return s
	}

	out := s[:1]
	for _, str := range s[1:] {
		if str != out[len(out)-1] {
			out = append(out, str)
		}
	}

	return out
}

func removeDuplicates2(s []string) []string {
	if len(s) <= 1 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}

	return s[:i+1]
}

func squashSpaces(arr []byte) []byte {
	if len(arr) <= 1 {
		return arr
	}

	i := 0
	for _, v := range arr[1:] {
		if !unicode.IsSpace(rune(v)) {
			i++
			arr[i] = v
		} else if !unicode.IsSpace(rune(arr[i])) {
			i++
			arr[i] = ' '
		}
	}
	return arr[:i+1]
}
