package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename2("a/b/c.go"))

	fmt.Println(basename("c.d.go"))
	fmt.Println(basename2("c.d.go"))

	fmt.Println(basename("abc"))
	fmt.Println(basename2("abc"))

	s := "vinicius"
	r := '/'
	fmt.Printf("s[0]: %v %[1]T\n", s[0])
	fmt.Printf("r: %v %[1]T\n", r)
}

// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
