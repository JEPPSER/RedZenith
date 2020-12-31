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

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
