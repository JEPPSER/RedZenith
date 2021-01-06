package common

// Direction ...
type Direction int

// Direction values ...
const (
	NONE Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

// Delta ...
var Delta float32

// Contains ...
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ClampF32 ...
func ClampF32(val float32, min float32, max float32) float32 {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

// Point ...
type Point struct {
	X float32
	Y float32
}

// IsInRect ...
func (p *Point) IsInRect(x float32, y float32, width float32, height float32) bool {
	return p.X >= x && p.X <= x+width && p.Y >= y && p.Y <= y+height
}
