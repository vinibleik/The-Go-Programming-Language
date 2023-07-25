package main

import (
	"fmt"
	"math"
)

// pc[i] is the populations count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the populations count (number of set bits) of x.
//
//	func popcount(x uint64) int {
//		return int(pc[byte(x>>(0*8))] +
//			pc[byte(x>>(1*8))] +
//			pc[byte(x>>(2*8))] +
//			pc[byte(x>>(3*8))] +
//			pc[byte(x>>(4*8))] +
//			pc[byte(x>>(5*8))] +
//			pc[byte(x>>(6*8))] +
//			pc[byte(x>>(7*8))])
//	}
func popcount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((x >> i) & 1)
	}
	// for x != 0 {
	// 	count += int(x & 1)
	// 	x >>= 1
	// }
	return count
}

func main() {
	var x uint64 = math.MaxUint64 >> 64
	fmt.Printf("PopCount(%064b) = %v\n", x, popcount(x))

}
