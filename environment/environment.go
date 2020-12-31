package environment

import (
	"red_zenith/common"
	"red_zenith/entities"
)

// Environment ...
type Environment struct {
	Gravity         float32
	MaxFallingSpeed float32
}

// Update ...
func (e Environment) Update(p *entities.Player) {

	// Gravity
	p.Y += p.YVelocity * common.Delta
	if p.YVelocity > e.MaxFallingSpeed {
		p.YVelocity = e.MaxFallingSpeed
	} else {
		p.YVelocity += e.Gravity * common.Delta
	}

	// Player movement
	p.X += p.XVelocity * common.Delta
}
