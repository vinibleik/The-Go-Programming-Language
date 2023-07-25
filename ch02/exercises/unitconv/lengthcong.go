package unitconv

import "fmt"

type (
	Feet   Length
	Meters Length
)

func (f Feet) String() string {
	return fmt.Sprintf("%.2gft", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%.2gm", m)
}

// FtToM converts Feet in Meters
func FtToM(f Feet) Meters {
	return Meters(f / 3.281)
}

func MToFt(m Meters) Feet {
	return Feet(m * 3.281)
}
