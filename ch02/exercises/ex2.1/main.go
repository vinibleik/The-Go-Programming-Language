package main

import (
	"fmt"

	"gopl.vini/ch02/tempconv"
)

func main() {

	fmt.Printf("tempconv.AbsoluteZeroK: %v\n", tempconv.AbsoluteZeroK)
	fmt.Printf("tempconv.FreezingK: %v\n", tempconv.FreezingK)
	fmt.Printf("tempconv.BoilingK: %v\n", tempconv.BoilingK)

	fmt.Printf("tempconv.KToC(273.15): %v\n", tempconv.KToC(273.15))
	fmt.Printf("tempconv.KToF(273.15): %v\n", tempconv.KToF(273.15))
}
