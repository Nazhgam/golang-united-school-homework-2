package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type Sides int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sides Sides) float64 {
	if sides == SidesTriangle {
		return (sideLen * sideLen * math.Sqrt(3)) / 4
	}
	if sides == SidesSquare {
		return sideLen * sideLen
	}
	if sides == SidesCircle {
		return math.Pi * sideLen * sideLen
	}
	return 0
}
