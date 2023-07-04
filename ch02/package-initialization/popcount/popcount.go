package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	total := 0
	for i := 0; i < 8; i++ {
		total += int(pc[byte(x>>(i*8))])
	}
	return total
}

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	total := 0
	for x != 0 {
		total += int(x & 1)
		x >>= 1
	}
	return total
}

// PopCountRightMost returns the population count (number of set bits) of x.
func PopCountRightMost(x uint64) int {
	// x & (x - 1) clears the rightmost non-zero bit of x
	i := 0
	for x != 0 {
		i++
		x = x & (x - 1)
	}
	return i
}
