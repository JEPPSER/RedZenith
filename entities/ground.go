package entities

import (
	"red_zenith/common"
)

// Ground ...
type Ground struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// OnCollision ...
func (g Ground) OnCollision(p *Player, dir common.Direction) {
	switch dir {
	case common.UP:
		p.YVelocity = 0
		p.Y = g.Y - p.Height - 1
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
