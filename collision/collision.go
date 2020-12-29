package collision

import (
	"math"
	"red_zenith/common"
	"red_zenith/entities"
)

// IsColliding ...
func IsColliding(p entities.Player, e entities.BaseEntity) bool {
	return p.X >= e.GetX()-p.Width &&
		p.X <= e.GetX()+e.GetWidth() &&
		p.Y >= e.GetY()-p.Height &&
		p.Y <= e.GetY()+e.GetHeight()
}

// GetCollisionDirection ...
func GetCollisionDirection(p entities.Player, e entities.BaseEntity) common.Direction {
	dir := common.NONE
	if IsColliding(p, e) {
		if p.Y >= e.GetY() && p.Y <= e.GetY()+e.GetHeight() {
			var deltaX float32 = 1000.0
			if p.X <= e.GetX() {
				deltaX = p.X + p.Width - e.GetX()
			} else if p.X+p.Width >= e.GetX()+e.GetWidth() {
				deltaX = p.X - (e.GetX() + e.GetWidth())
			}

			var deltaY float32 = e.GetY() + e.GetHeight() - p.Y
			if deltaY <= abs(deltaX) {
				dir = common.DOWN
			} else if deltaX > 0.0 {
				dir = common.LEFT
			} else if deltaX < 0.0 {
				dir = common.RIGHT
			}
		} else if p.Y+p.Height >= e.GetY() && p.Y+p.Height <= e.GetY()+e.GetHeight() {
			var deltaX float32 = 1000.0
			if p.X <= e.GetX() {
				deltaX = p.X + p.Width - e.GetX()
			} else if p.X+p.Width >= e.GetX()+e.GetWidth() {
				deltaX = p.X - (e.GetX() + e.GetWidth())
			}

			var deltaY = p.Y + p.Height - e.GetY()
			if deltaY <= abs(deltaX) {
				dir = common.UP
			} else if deltaX > 0.0 {
				dir = common.LEFT
			} else if deltaX < 0.0 {
				dir = common.RIGHT
			}
		} else {
			if p.X < e.GetX() {
				dir = common.LEFT
			} else if p.X+p.Width > e.GetX()+e.GetWidth() {
				dir = common.RIGHT
			}
		}
	}
	return dir
}

func abs(val float32) float32 {
	return float32(math.Abs(float64(val)))
}
