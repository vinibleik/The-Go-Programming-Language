package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", comma("12345"))
	fmt.Printf("%s\n", comma("125"))
	fmt.Printf("%s\n", comma("1234522420"))
}

// comma insert commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
