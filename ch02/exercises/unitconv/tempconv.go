package unitconv

import "fmt"

type (
	Celsius    Temperature
	Fahrenheit Temperature
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2gºF", f)
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
