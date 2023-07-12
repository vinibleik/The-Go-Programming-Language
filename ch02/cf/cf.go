// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.vini/ch02/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s = %s\n",
			k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
