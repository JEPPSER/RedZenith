package controller

import (
	"github.com/veandco/go-sdl2/sdl"
	"red_zenith/common"
	"red_zenith/entities"
)

// PlayerController ...
type PlayerController struct {
}

const jumpKey = sdl.SCANCODE_SPACE
const leftKey = sdl.SCANCODE_LEFT
const rightKey = sdl.SCANCODE_RIGHT

const playerAcc = 0.01
const playerSpeed = 0.7
const playerJumpForce = 2

// Control ...
func (controller *PlayerController) Control(p *entities.Player, input []int) {

	// Player movement controls
	if !p.IsGrounded {
		p.CanJump = false
	} else if !p.CanJump && !common.Contains(input, jumpKey) {
		p.CanJump = true
	}
	if common.Contains(input, rightKey) {
		move(p, 1)
	}
	if common.Contains(input, leftKey) {
		move(p, -1)
	}
	if !common.Contains(input, rightKey) && !common.Contains(input, leftKey) {
		stop(p)
	}
	if common.Contains(input, jumpKey) && p.CanJump {
		p.YVelocity = -playerJumpForce
		p.CanJump = false
	}

	// Player item controls
	p.Item.Update(input, p)
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
