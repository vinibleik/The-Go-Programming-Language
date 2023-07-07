package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	bytes := []byte(s)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("bytes: %v\n", bytes)
	s = "hello, worldź"
	bytes = []byte(s)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("bytes: %v\n", bytes)
	fmt.Printf("s[len(s)-1]: %v %[1]c\n", s[len(s)-1])

	fmt.Printf("s[0:5]: %v\n", s[0:5])
	fmt.Printf("s[:5]: %v\n", s[:5])
	fmt.Printf("s[7:]: %v\n", s[7:])
	fmt.Printf("s[:]: %v\n", s[:])

	a := "left foot"
	t := a
	left := a[:4]
	foot := a[5:]
	a += ", right foot"

	fmt.Println(a) // "left foot, right foot"
	fmt.Println(t) // "left foot"
	fmt.Println(left)
	fmt.Println(foot)

	t = "vinicius"

	fmt.Println(a) // "left foot, right foot"
	fmt.Println(t) // "left foot"
	fmt.Println(left)
	fmt.Println(foot)

	fmt.Println("\a")

	fmt.Println(`\a`)

	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
		go command [arguments]
	...
`
	fmt.Printf("%s", GoUsage)

	var unicode_str = "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Printf("unicode_str: %v\n", unicode_str)
	unicode_str = "\u4e16\u754c"
	fmt.Printf("unicode_str: %v\n", unicode_str)
	unicode_str = "\U00004e16\U0000754c"
	fmt.Printf("unicode_str: %v\n", unicode_str)

	s = "Hello, 世界"
	fmt.Printf("s: %v\n", s)
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("utf8.RuneCountInString(s): %v\n", utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("i = %d\tr = %c\tsize = %v\n", i, r, size)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}

	fmt.Printf("n: %v\n", n)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]: %[2]v \t %[2]q \t %[2]d \t %[2]c \n", i, s[i])
	}

	fmt.Printf("%c\n", '\uFFFD')

	s = "プログラム"
	fmt.Printf("s: % x\n", s)
	r := []rune(s)
	fmt.Printf("r: % x\n", r)

	fmt.Printf("s: %v\n", s)
	fmt.Printf("r: %v\n", r)
	fmt.Printf("string(r): %v\n", string(r))

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
