package entities

import (
	"github.com/veandco/go-sdl2/sdl"
	"red_zenith/common"
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
func (p *Player) Render(renderer *sdl.Renderer, offsetX float32, offsetY float32) {
	p.Item.Render(renderer, offsetX, offsetY)

	src := &sdl.Rect{X: 0, Y: 0, W: 32, H: 32}
	target := &sdl.Rect{X: int32(p.X - offsetX), Y: int32(p.Y - offsetY), W: int32(p.Width), H: int32(p.Height)}
	angle := p.X
	renderer.CopyEx(common.PlayerImage, src, target, float64(angle), &sdl.Point{X: 16, Y: 16}, sdl.FLIP_NONE)
	renderer.Copy(common.PlayerOverlayImage, src, target)
}
