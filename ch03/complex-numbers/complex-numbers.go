package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	y := 3 + 4i
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(.1i)
}
