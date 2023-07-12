package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("For loop")
	start := time.Now()

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%v microseconds elapsed%c%c", time.Since(start).Microseconds(), 10, 10)

	fmt.Println("For loop")
	start = time.Now()

	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%2v microseconds elapsed%c", time.Since(start).Microseconds(), 10)

}
