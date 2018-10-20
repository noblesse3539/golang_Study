package lenconv

import (
	"fmt"
)

type Meter float64
type Feet float64

const ()

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }

func FtToM(ft Feet) Meter {
	return Meter(ft * 0.3048)
}
func MToFt(m Meter) Feet {
	return Feet(m * 3.2808)
}
