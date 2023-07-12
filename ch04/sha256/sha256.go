package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	var x byte = 0b11110000
	var y byte = 0b10101010

	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)
	fmt.Printf("(x ^ y): %08b\n", (x ^ y))
	fmt.Printf("PopCount(x ^ y): %v\n", PopCount(x^y))

	fmt.Printf("differenceBits(c1, c2): %v\n", differenceBits(c1, c2))

	for i := range c1 {
		c1[i] = x
	}

	for i := range c2 {
		c2[i] = y
	}
	fmt.Printf("differenceBits(c1, c2): %v\n", differenceBits(c1, c2))
}

func differenceBits(a [32]byte, b [32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		diff += PopCount(a[i] ^ b[i])
	}
	return diff
}

func PopCount(x byte) int {
	// x & (x - 1) clears the rightmost non-zero bit of x
	i := 0
	for x != 0 {
		i++
		x &= (x - 1)
	}
	return i
}
