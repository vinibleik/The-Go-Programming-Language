package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Time for inefficient approach: ", time.Since(start).Microseconds())

	start = time.Now()
	s = strings.Join(os.Args, " ")
	fmt.Println("Time for strings.Join approach: ", time.Since(start).Microseconds())
}
