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

const useKey = sdl.SCANCODE_D
const aimRightKey = sdl.SCANCODE_S
const aimLeftKey = sdl.SCANCODE_A

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
		if common.Contains(input, aimLeftKey) {
			h.Angle += 0.005 * common.Delta
		} else if common.Contains(input, aimRightKey) {
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
		h.EndPoint.X = x + (50+h.shotDuration*1.5)*float32(math.Sin(float64(h.Angle)))
		h.EndPoint.Y = y + (50+h.shotDuration*1.5)*float32(math.Cos(float64(h.Angle)))

		if h.isConnected() {
			p.XVelocity = float32(math.Sin(float64(h.Angle))) * 3
			p.YVelocity = float32(math.Cos(float64(h.Angle))) * 3
			h.isShooting = false
			h.shotDuration = 0
		} else {
			if h.shotDuration > 300 {
				h.isShooting = false
				h.shotDuration = 0
			} else {
				h.shotDuration += common.Delta
			}
		}
	}
}

// Render ...
func (h HookShot) Render(renderer *sdl.Renderer, offsetX float32, offsetY float32) {
	//renderer.SetDrawColor(0, 255, 0, 255)
	//renderer.DrawLine(int32(h.StartPoint.X-offsetX), int32(h.StartPoint.Y-offsetY), int32(h.EndPoint.X-offsetX), int32(h.EndPoint.Y-offsetY))

	angle := float64(h.Angle*(180/math.Pi)) * -1

	// Head
	src := &sdl.Rect{X: 0, Y: 0, W: 16, H: 16}
	target := &sdl.Rect{X: int32(h.EndPoint.X - 8 - offsetX), Y: int32(h.EndPoint.Y - offsetY), W: 16, H: 16}
	renderer.CopyEx(common.HookshotHead, src, target, angle, &sdl.Point{X: 8, Y: 0}, sdl.FLIP_NONE)

	// Chain
	src = &sdl.Rect{X: 0, Y: 0, W: 6, H: 12}
	full := 40 + h.shotDuration*1.5
	for i := 12; i < int(full); i += 12 {
		target = &sdl.Rect{X: int32(h.EndPoint.X - 3 - offsetX), Y: int32(h.EndPoint.Y - float32(i) - offsetY), W: 6, H: 12}
		renderer.CopyEx(common.HookshotChain, src, target, angle, &sdl.Point{X: 3, Y: int32(i)}, sdl.FLIP_NONE)
	}
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
