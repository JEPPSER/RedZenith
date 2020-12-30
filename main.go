package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"red_zenith/collision"
	"red_zenith/common"
	"red_zenith/entities"
	"strconv"
)

var width int32
var height int32

func main() {
	width = 1000
	height = 700

	canJump := false

	objects := []entities.BaseEntity{}

	player := &entities.Player{
		X:          300,
		Y:          400,
		Width:      20,
		Height:     20,
		YVelocity:  0,
		XVelocity:  0,
		IsGrounded: false,
	}

	e1 := &entities.Ground{
		X:      200,
		Y:      610,
		Width:  600,
		Height: 50,
	}

	e2 := &entities.Ground{
		X:      500,
		Y:      350,
		Width:  50,
		Height: 200,
	}

	e3 := &entities.Ground{
		X:      700,
		Y:      500,
		Width:  50,
		Height: 100,
	}

	objects = append(objects, e1)
	objects = append(objects, e2)
	objects = append(objects, e3)

	window, renderer := initSDL()
	defer window.Destroy()
	defer renderer.Destroy()

	buffer, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, width, height)
	if err != nil {
		panic(err)
	}
	defer buffer.Destroy()

	input := []int{}
	time := sdl.GetTicks()
	count := 0
	secondTimer := float32(0)

	// Main loop
	for {
		// Update timing
		newTime := sdl.GetTicks()
		common.Delta = float32(newTime - time)
		time = newTime
		secondTimer += common.Delta

		// Input
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				input = keyboardPressed(event.(*sdl.KeyboardEvent), input)
			}
		}

		// Controls
		if !canJump && player.IsGrounded && !contains(input, sdl.SCANCODE_SPACE) {
			canJump = true
		}
		if contains(input, sdl.SCANCODE_RIGHT) {
			player.X += 0.7 * common.Delta
		}
		if contains(input, sdl.SCANCODE_LEFT) {
			player.X -= 0.7 * common.Delta
		}
		if contains(input, sdl.SCANCODE_SPACE) && canJump {
			player.YVelocity = -2
			canJump = false
		}

		// Physic
		player.Y += player.YVelocity * common.Delta
		player.YVelocity += 0.01 * common.Delta

		// Collision
		player.IsGrounded = false
		for _, e := range objects {
			dir := collision.GetCollisionDirection(*player, e)
			e.OnCollision(player, dir)
		}

		// Render
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 0, 255)
		renderer.FillRect(&sdl.Rect{X: int32(player.X), Y: int32(player.Y), W: int32(player.Width), H: int32(player.Height)})
		renderer.SetDrawColor(255, 0, 0, 255)
		for _, e := range objects {
			renderer.FillRect(&sdl.Rect{X: int32(e.GetX()), Y: int32(e.GetY()), W: int32(e.GetWidth()), H: int32(e.GetHeight())})
		}
		renderer.Present()

		// FPS counter
		count++
		if secondTimer > 1000 {
			window.SetTitle("GoGame  FPS: " + strconv.FormatInt(int64(count), 10))
			count = 0
			secondTimer = 0
		}
	}
}

func keyboardPressed(e *sdl.KeyboardEvent, input []int) []int {
	var scancode = int(e.Keysym.Scancode)

	if e.Type == sdl.KEYDOWN && !contains(input, scancode) {
		input = append(input, scancode)
	} else if e.Type == sdl.KEYUP && contains(input, scancode) {
		input = remove(input, scancode)
	}

	return input
}

func initSDL() (*sdl.Window, *sdl.Renderer) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("RedZenith", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	window.SetResizable(false)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	return window, renderer
}

func remove(s []int, e int) []int {
	var index = indexOf(s, e)
	s[index] = s[len(s)-1]
	s[len(s)-1] = 0
	s = s[:len(s)-1]
	return s
}

func indexOf(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
