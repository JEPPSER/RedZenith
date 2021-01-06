package camera

import (
	"github.com/veandco/go-sdl2/sdl"
	"red_zenith/entities"
)

var yScrollValue float32 = 150.0

// Camera ...
type Camera struct {
	OffsetX float32
	OffsetY float32
	Player  *entities.Player
	Objects *[]entities.BaseEntity
}

// Render ...
func (c *Camera) Render(renderer *sdl.Renderer) {
	// Set camera position
	c.OffsetX = c.Player.X - float32(renderer.GetViewport().W)/2.0

	if c.Player.Y-c.OffsetY < yScrollValue {
		c.OffsetY = c.Player.Y - yScrollValue
	} else if c.Player.Y-c.OffsetY > float32(renderer.GetViewport().H)-yScrollValue-c.Player.Height {
		c.OffsetY = c.Player.Y - (float32(renderer.GetViewport().H) - yScrollValue - c.Player.Height)
	}

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	for _, e := range *c.Objects {
		e.Render(renderer, c.OffsetX, c.OffsetY)
	}
	c.Player.Render(renderer, c.OffsetX, c.OffsetY)
	renderer.Present()
}
