package entities

import (
	"github.com/veandco/go-sdl2/sdl"
	"red_zenith/common"
)

// BaseEntity ...
type BaseEntity interface {
	OnCollision(p *Player, dir common.Direction)
	GetX() float32
	GetY() float32
	GetWidth() float32
	GetHeight() float32
	Render(render *sdl.Renderer, offsetX float32, offsetY float32)
}
