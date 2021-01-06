package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"red_zenith/camera"
	"red_zenith/collision"
	"red_zenith/common"
	"red_zenith/controller"
	"red_zenith/entities"
	"red_zenith/environment"
	"red_zenith/level"
	"strconv"
)

var width int32
var height int32

func main() {
	width = 1280
	height = 720

	objects, spawnPoint := level.GetLevel("level.rz")

	hookShot := &entities.HookShot{
		EndPoint:   common.Point{X: 0, Y: 0},
		StartPoint: common.Point{X: 0, Y: 0},
		Angle:      3.14,
		Objects:    &objects,
	}
	player := &entities.Player{
		X:          spawnPoint.X,
		Y:          spawnPoint.Y,
		Width:      32,
		Height:     32,
		YVelocity:  0,
		XVelocity:  0,
		IsGrounded: false,
		CanJump:    false,
		Item:       hookShot,
	}

	camera := &camera.Camera{
		OffsetX: 0,
		OffsetY: 0,
		Player:  player,
		Objects: &objects,
	}

	controller := &controller.PlayerController{}

	environment := &environment.Environment{
		Gravity:         0.01,
		MaxFallingSpeed: 1.5,
	}

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
		controller.Control(player, input)

		// Physics
		environment.Update(player)

		// Collision
		player.IsGrounded = false
		for _, e := range objects {
			dir := collision.GetCollisionDirection(*player, e)
			e.OnCollision(player, dir)
		}

		// Render
		camera.Render(renderer)

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

	if e.Type == sdl.KEYDOWN && !common.Contains(input, scancode) {
		input = append(input, scancode)
	} else if e.Type == sdl.KEYUP && common.Contains(input, scancode) {
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
