package entities

import (
	"red_zenith/common"
)

// Ground ...
type Ground struct {
	X               float32
	Y               float32
	Width           float32
	Height          float32
	CollisionTimers []float32
}

// OnCollision ...
func (g *Ground) OnCollision(p *Player, dir common.Direction) {

	// Reset collision timers
	if len(g.CollisionTimers) == 0 || dir == common.NONE {
		g.ClearCollisionTimers()
	}

	// Count current collision time
	if dir != common.NONE {
		g.CollisionTimers[dir-1] += common.Delta
	}

	if dir != common.NONE && g.CollisionTimers[dir-1] > 5 {
		switch dir {
		case common.UP:
			p.YVelocity = 0
			p.Y = g.Y - p.Height
			p.IsGrounded = true
		case common.DOWN:
			p.YVelocity = 0
			p.Y = g.Y + g.Height + 1
		case common.LEFT:
			p.XVelocity = 0
			p.X = g.X - p.Width
		case common.RIGHT:
			p.XVelocity = 0
			p.X = g.X + g.Width
		default:
		}
	}
}

// ClearCollisionTimers ...
func (g *Ground) ClearCollisionTimers() {
	g.CollisionTimers = []float32{0.0, 0.0, 0.0, 0.0}
}

// GetX ...
func (g Ground) GetX() float32 {
	return g.X
}

// GetY ...
func (g Ground) GetY() float32 {
	return g.Y
}

// GetWidth ...
func (g Ground) GetWidth() float32 {
	return g.Width
}

// GetHeight ...
func (g Ground) GetHeight() float32 {
	return g.Height
}
