package entities

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"red_zenith/common"
)

// UsableItem ...
type UsableItem interface {
	Update(input []int, p *Player)
	Render(renderer *sdl.Renderer, offsetX float32, offsetY float32)
}

const useKey = sdl.SCANCODE_A
const aimRightKey = sdl.SCANCODE_S
const aimLeftKey = sdl.SCANCODE_D

// HookShot ...
type HookShot struct {
	isShooting   bool
	shotDuration float32
	canShoot     bool
	shotCooldown float32

	Angle      float32
	EndPoint   common.Point
	StartPoint common.Point
	Objects    *[]BaseEntity
}

// Update ...
func (h *HookShot) Update(input []int, p *Player) {
	x := p.X + p.Width/2.0
	y := p.Y + p.Height/2.0
	h.StartPoint.X = x
	h.StartPoint.Y = y

	if !h.canShoot {
		if !common.Contains(input, useKey) && h.shotCooldown > 1000 {
			h.canShoot = true
			h.shotCooldown = 0
		}
		h.shotCooldown += common.Delta
	}

	if !h.isShooting {
		if common.Contains(input, aimRightKey) {
			h.Angle += 0.005 * common.Delta
		} else if common.Contains(input, aimLeftKey) {
			h.Angle -= 0.005 * common.Delta
		}

		h.EndPoint.X = x + float32(50*math.Sin(float64(h.Angle)))
		h.EndPoint.Y = y + float32(50*math.Cos(float64(h.Angle)))

		if h.canShoot && common.Contains(input, useKey) {
			h.isShooting = true
			h.canShoot = false
			h.shotDuration = 0
		}
	} else {
		h.EndPoint.X = x + h.shotDuration*float32(math.Sin(float64(h.Angle)))
		h.EndPoint.Y = y + h.shotDuration*float32(math.Cos(float64(h.Angle)))

		if h.isConnected() {
			p.XVelocity = float32(math.Sin(float64(h.Angle))) * 3
			p.YVelocity = float32(math.Cos(float64(h.Angle))) * 3
			h.isShooting = false
			h.shotDuration = 0
		} else {
			if h.shotDuration > 500 {
				h.isShooting = false
			} else {
				h.shotDuration += common.Delta
			}
		}
	}
}

// Render ...
func (h HookShot) Render(renderer *sdl.Renderer, offsetX float32, offsetY float32) {
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLine(int32(h.StartPoint.X-offsetX), int32(h.StartPoint.Y-offsetY), int32(h.EndPoint.X-offsetX), int32(h.EndPoint.Y-offsetY))
}

func (h HookShot) isConnected() bool {
	for _, o := range *h.Objects {
		ground, ok := o.(*Ground)
		if ok {
			if h.EndPoint.IsInRect(ground.X, ground.Y, ground.Width, ground.Height) {
				return true
			}
		}
	}
	return false
}
