package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.vini/ch02/exercises/unitconv"
)

func printWeight(n float64) {
	p := unitconv.Pounds(n)
	k := unitconv.Kilograms(n)
	fmt.Printf("%s = %s\n%s = %s\n", k, unitconv.KgToP(k), p, unitconv.PToKg(p))
}

func printLength(n float64) {
	ft := unitconv.Feet(n)
	m := unitconv.Meters(n)
	fmt.Printf("%s = %s\n%s = %s\n", m, unitconv.MToFt(m), ft, unitconv.FtToM(ft))
}

func printTemp(n float64) {
	c := unitconv.Celsius(n)
	f := unitconv.Fahrenheit(n)
	fmt.Printf("%s = %s\n%s = %s\n", c, unitconv.CToF(c), f, unitconv.FToC(f))
}

func printAll(n float64) {
	fmt.Printf("Temperatures %.2g\n", n)
	printTemp(n)
	fmt.Printf("\nWeight %.2g\n", n)
	printWeight(n)
	fmt.Printf("\nLengths %.2g\n", n)
	printLength(n)
	fmt.Println()
}

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			v := input.Text()
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Bad float number: %v\n\n", v)
				continue
			}
			printAll(n)
		}
	} else {
		for _, v := range os.Args[1:] {
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Bad float number: %v\n\n", v)
				continue
			}
			printAll(n)
		}
	}
}
