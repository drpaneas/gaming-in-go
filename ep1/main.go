package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize all the sub-systems
	// INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS | INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL failed:", err)
		return
	}

	// Declare a window (returns the window that was created or an error)
	window, err := sdl.CreateWindow(
		"Gaming episode 2",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		600,
		800,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	// Make sure Window gets destroyes when main() returns
	defer window.Destroy()

	// Create a Renderer (a layer that exists ontop of the window)
	// It is sth we can use to draw
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
		return
	}
	// Make sure Renderer gets destroyed when we're done with it
	defer renderer.Destroy()
}
