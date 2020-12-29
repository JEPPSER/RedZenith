package entities

import "red_zenith/common"

// BaseEntity ...
type BaseEntity interface {
	OnCollision(p *Player, dir common.Direction)
	GetX() float32
	GetY() float32
	GetWidth() float32
	GetHeight() float32
}
