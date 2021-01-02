package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Player ...
type Player struct {
	X          float32
	Y          float32
	Width      float32
	Height     float32
	XVelocity  float32
	YVelocity  float32
	IsGrounded bool
	CanJump    bool
	Item       UsableItem
}

// Render ...
func (p *Player) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.FillRect(&sdl.Rect{X: int32(p.X), Y: int32(p.Y), W: int32(p.Width), H: int32(p.Height)})
	p.Item.Render(renderer)
}
