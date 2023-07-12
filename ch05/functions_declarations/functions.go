package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func hypot(p Point) (hyp float64) {
	hyp = math.Sqrt(p.X*p.X + p.Y*p.Y)
	return
}

func main() {
	fmt.Println(hypot(Point{Y: 4, X: 3}))
	fmt.Printf("%T\n", hypot)
}