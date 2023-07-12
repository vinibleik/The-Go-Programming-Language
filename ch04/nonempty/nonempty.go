// Nonempty is an example of an in-place slice algorithm
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q len(data): %d cap(data): %d\n", data, len(data), cap(data))
	newData := nonempty(data)
	fmt.Printf("%q len(newData): %d cap(newData): %d\n", newData, len(newData), cap(newData))
	fmt.Printf("%q len(data): %d cap(data): %d\n", data, len(data), cap(data))

	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("s: %v\n", s)
	fmt.Printf("remove(s, 2): %v\n", remove(s, 2))
	fmt.Printf("s: %v\n", s)
}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
