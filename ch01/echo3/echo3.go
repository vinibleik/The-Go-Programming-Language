package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(args)
}
