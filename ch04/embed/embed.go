package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type newCircle struct {
	Point
	Radius int
}

type newWheel struct {
	newCircle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Printf("w: %v\n", w)

	w.Circle.Center = Point{Y: 5}
	fmt.Printf("w: %v\n", w)

	w.Circle = Circle{Center: Point{X: 4, Y: 0}}
	fmt.Printf("w: %v\n", w)

	var x newWheel
	x.X = 8			// Equivalent to x.Circle.Point.X = 8
	x.Y = 8			// Equivalent to x.Circle.Point.Y = 8
	x.Radius = 4	// Equivalent to x.Circle.Radius = 4
	x.Spokes = 20 
	fmt.Printf("x: %#v\n", x)
}