package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	x := FToC(212.0)
	fmt.Println(x.String())
	fmt.Printf("%v\n", x)
	fmt.Printf("%s\n", x)
	fmt.Println(x)
	fmt.Printf("%g\n", x)
	fmt.Println(float64(x))
}
