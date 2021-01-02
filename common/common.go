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

// Point ...
type Point struct {
	X float32
	Y float32
}

// IsInRect ...
func (p *Point) IsInRect(x float32, y float32, width float32, height float32) bool {
	return p.X >= x && p.X <= x+width && p.Y >= y && p.Y <= y+height
}
