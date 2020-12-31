package controller

import (
	"github.com/veandco/go-sdl2/sdl"
	"red_zenith/common"
	"red_zenith/entities"
)

// PlayerController ...
type PlayerController struct {
}

const playerAcc = 0.01
const playerSpeed = 0.7
const playerJumpForce = 2

// Control ...
func (controller *PlayerController) Control(p *entities.Player, input []int) {
	if !p.IsGrounded {
		p.CanJump = false
	} else if !p.CanJump && !common.Contains(input, sdl.SCANCODE_SPACE) {
		p.CanJump = true
	}
	if common.Contains(input, sdl.SCANCODE_RIGHT) {
		move(p, 1)
	}
	if common.Contains(input, sdl.SCANCODE_LEFT) {
		move(p, -1)
	}
	if !common.Contains(input, sdl.SCANCODE_RIGHT) && !common.Contains(input, sdl.SCANCODE_LEFT) {
		stop(p)
	}
	if common.Contains(input, sdl.SCANCODE_SPACE) && p.CanJump {
		p.YVelocity = -playerJumpForce
		p.CanJump = false
	}
}

func stop(p *entities.Player) {
	if p.XVelocity < playerAcc && p.XVelocity > -playerAcc {
		p.XVelocity = 0
		return
	} else if p.XVelocity > 0 {
		p.XVelocity -= playerAcc * common.Delta
	} else if p.XVelocity < 0 {
		p.XVelocity += playerAcc * common.Delta
	}
}

func move(p *entities.Player, dir int) {
	if p.XVelocity*float32(dir) < playerSpeed {
		p.XVelocity += playerAcc * float32(dir) * common.Delta
	}
}
