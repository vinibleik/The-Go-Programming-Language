package main

import (
	"fmt"
	"math"

	"gopl.vini/ch02/popcount"
)

func main() {
	var x uint64 = math.MaxUint64 - 34214311
	fmt.Printf("PopCount(%b) = %v\n", x, popcount.PopCount(x))
	fmt.Printf("PopCountLoop(%b) = %v\n", x, popcount.PopCountLoop(x))
	fmt.Printf("PopCountShift(%b) = %v\n", x, popcount.PopCountShift(x))
	fmt.Printf("PopCountRightMost(%b) = %v\n", x, popcount.PopCountRightMost(x))

	// fmt.Printf("%b\n%b\n", x, x&(x-1))

	// for i := 0; i < 256; i++ {
	// 	fmt.Printf("i: %v i/2 = %v\n", i, i/2)
	// }
}
