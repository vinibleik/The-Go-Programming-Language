package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type Point struct {
	X, Y int
}

func main() {
	var dilbert Employee
	fmt.Printf("dilbert: %v\n", dilbert)
	dilbert.Salary -= 5000
	fmt.Printf("dilbert: %v\n", dilbert)
	position := &dilbert.Position
	*position = "Senior" + *position
	fmt.Printf("dilbert: %v\n", dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Name = "Dilbert"
	fmt.Printf("employeeOfTheMonth: %v\n", employeeOfTheMonth)
	fmt.Printf("(*employeeOfTheMonth): %v\n", (*employeeOfTheMonth))
	fmt.Printf("dilbert: %v\n", dilbert)

	p := Point{1, 2}
	fmt.Printf("p: %v\n", p)
	x := Point{X: 1, Y: 2}
	fmt.Printf("x: %v\n", x)
	x = Point{Y: 2, X: 1}
	fmt.Printf("x: %v\n", x)
	x = Point{Y: 2}
	fmt.Printf("x: %v\n", x)
	x = Point{X: 1}
	fmt.Printf("x: %v\n", x)

	fmt.Println("")
	fmt.Printf("p: %v\n", p)
	z := Scale(p,5)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("z: %v\n", z)

	InScale(&p, 3)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("InScale(&Point{3, 4}, 10): %v\n", *InScale(&Point{3, 4}, 10))

	fmt.Println(p.X == x.X && p.Y == x.Y)
	fmt.Println(p == x)
	x = Point{3, 6}
	fmt.Println(p.X == x.X && p.Y == x.Y)
	fmt.Println(p == x)
	
	type address struct {
		hostname string
		port int
	}

	hits := make(map[address]int)
	fmt.Printf("hits: %v\n", hits)
	hits[address{"golang.org", 443}]++
	fmt.Printf("hits: %v\n", hits)
}

func Scale(p Point, factor int) Point {
	p.X *= factor
	p.Y *= factor
	return p
	// return Point{p.X * factor, p.Y * factor}
}

func InScale(p *Point, factor int) *Point{
	p.X *= factor
	(*p).Y *= factor
	return p
}