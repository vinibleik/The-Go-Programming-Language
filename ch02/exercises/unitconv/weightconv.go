package unitconv

import "fmt"

type (
	Pounds    Weight
	Kilograms Weight
)

func (p Pounds) String() string {
	return fmt.Sprintf("%.2glb", p)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%.2gKg", k)
}

// PToKg converts Pounds in Kilograms
func PToKg(p Pounds) Kilograms {
	return Kilograms(p / 2.205)
}

// KgToP converts Kilograms in Pounds
func KgToP(k Kilograms) Pounds {
	return Pounds(k * 2.205)
}
