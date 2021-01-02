package entities

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"red_zenith/common"
)

// UsableItem ...
type UsableItem interface {
	Update(input []int, p *Player)
	Render(renderer *sdl.Renderer)
}

const useKey = sdl.SCANCODE_A
const aimRightKey = sdl.SCANCODE_S
const aimLeftKey = sdl.SCANCODE_D

// HookShot ...
type HookShot struct {
	isShooting   bool
	shotDuration float32

	Angle      float32
	EndPoint   common.Point
	StartPoint common.Point
	Objects    *[]BaseEntity
}

// Update ...
func (h *HookShot) Update(input []int, p *Player) {
	x := p.X + p.Width/2.0
	y := p.Y + p.Height/2.0
	if !h.isShooting {
		if common.Contains(input, aimRightKey) {
			h.Angle += 0.005 * common.Delta
		} else if common.Contains(input, aimLeftKey) {
			h.Angle -= 0.005 * common.Delta
		}
	}

	h.StartPoint.X = x
	h.StartPoint.Y = y
	if !h.isShooting {
		h.EndPoint.X = x + float32(50*math.Sin(float64(h.Angle)))
		h.EndPoint.Y = y + float32(50*math.Cos(float64(h.Angle)))
	}
	if !h.isShooting && common.Contains(input, useKey) {
		h.isShooting = true
		h.shotDuration = 0
	}

	if h.isShooting {
		h.EndPoint.X = x + h.shotDuration*float32(math.Sin(float64(h.Angle)))
		h.EndPoint.Y = y + h.shotDuration*float32(math.Cos(float64(h.Angle)))

		if h.isConnected() {
			p.XVelocity = float32(math.Sin(float64(h.Angle))) * 3
			p.YVelocity = float32(math.Cos(float64(h.Angle))) * 3
			h.isShooting = false
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
func (h HookShot) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLine(int32(h.StartPoint.X), int32(h.StartPoint.Y), int32(h.EndPoint.X), int32(h.EndPoint.Y))
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
