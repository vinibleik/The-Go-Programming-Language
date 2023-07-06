package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.MaxFloat32: %v\n", math.MaxFloat32)
	fmt.Printf("math.MaxFloat64: %v\n", math.MaxFloat64)

	fmt.Printf("math.MaxFloat32: %g\n", math.MaxFloat32)
	fmt.Printf("math.MaxFloat64: %g\n", math.MaxFloat64)

	var f float32 = 1 << 24
	fmt.Printf("f: %v f + 1 = %v\n", f, f+1)
	fmt.Println(f == f+1)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(math.IsNaN(nan), nan, math.NaN())
	fmt.Println(nan == nan, nan < nan, nan > nan, nan <= nan, nan >= nan, nan != nan)
	fmt.Println(nan == math.NaN(), nan != math.NaN())
}
