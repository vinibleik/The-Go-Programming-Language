package main

import (
	"fmt"
	"time"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB = 1000
	MB = KB * KB
	GB = KB * MB
	TB = KB * GB
	PB = KB * TB
	EB = KB * PB
	ZB = KB * EB
	YB = KB * ZB
)

func main() {
	const (
		e  = 2.71828182845904523536028747135266249775724709369995957496696763
		pi = 3.14159265358979323846264338327950288419716939937510582097494459
	)

	fmt.Printf("e: %v\n", e)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d)

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Printf("%3d -> %05[1]b\n", FlagUp)
	fmt.Printf("%3d -> %05[1]b\n", FlagBroadcast)
	fmt.Printf("%3d -> %05[1]b\n", FlagLoopback)
	fmt.Printf("%3d -> %05[1]b\n", FlagPointToPoint)
	fmt.Printf("%3d -> %05[1]b\n", FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"

	var z uint8 = 0xff
	var y uint8 = 0b00100110

	fmt.Printf("%8s: %08b\n", "z", z)
	fmt.Printf("%8s: %08b\n", "y", y)
	fmt.Printf("%8s: %08b\n", "^y", ^y)
	fmt.Printf("(z &^ y): %08b\n", (z &^ y))

	fmt.Printf("KiB: %v\n", KiB)
	fmt.Printf("MiB: %v\n", MiB)
	fmt.Printf("GiB: %v\n", GiB)
	fmt.Printf("TiB: %v\n", TiB)
	fmt.Printf("PiB: %v\n", PiB)
	fmt.Printf("EiB: %v\n", EiB)
	// fmt.Printf("ZiB: %v\n", ZiB)
	// fmt.Printf("YiB: %v\n", YiB)

	fmt.Printf("KB: %v\n", KB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)
	fmt.Printf("TB: %v\n", TB)
	fmt.Printf("PB: %v\n", PB)
	fmt.Printf("EB: %v\n", EB)
	// fmt.Printf("ZB: %g\n", ZB)
	// fmt.Printf("YB: %v\n", YB)

	fmt.Println(YiB / ZiB)

	const i = 1
	const j int = 1

	fmt.Println(1.3 / i)
	// fmt.Println(1.3 / j)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	const (
		deadbeef = 0xdeadbeef        // untyped int with value 3735928559
		a2       = uint32(deadbeef)  // uint32 with value 3735928559
		b2       = float32(deadbeef) // float32 with value 3735928576 (rounded up)
		c2       = float64(deadbeef) // float64 with value 3735928559 (exact)
		// d2       = int32(deadbeef)   // compile error: constant overflows int32
		// e2       = float64(1e309)    // compile error: constant overflows float64
		// f2       = uint(-1)          // compile error: constant underflows uint
	)
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}
