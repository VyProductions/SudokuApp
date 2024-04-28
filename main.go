package main

import (
	"log"
	"time"

	"main/engine"
	"main/selection"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Error initializing SDL: %s\n", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Error initializing TTF: %s\n", err)
	}
	defer ttf.Quit()

	wind_width, wind_height := int32(800), int32(600)
	window, err := sdl.CreateWindow("Sudoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, wind_width, wind_height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Error creating window: %s\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Error creating renderer: %s\n", err)
	}
	defer renderer.Destroy()

	appEngine := &engine.Engine{}

	err = appEngine.Setup(window, renderer)

	if err != nil {
		log.Fatalf("Error starting engine: %s\n", err)
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				appEngine.MoveMouse(sdl.Point{X: t.X, Y: t.Y})
			case *sdl.MouseButtonEvent:
				press_state, _ := selection.Ternary(event.GetType() == sdl.MOUSEBUTTONDOWN, engine.PRESSED, engine.RELEASED).(byte)
				args := []interface{}{t.X, t.Y}
				appEngine.ProcessAction(appEngine.InputTransform[int(t.Button)], press_state, args)
			case *sdl.MouseWheelEvent:
				action_byte, _ := selection.Ternary(t.X == 0, engine.VERT_SCROLL, engine.HORIZ_SCROLL).(byte)
				press_state, _ := selection.Ternary(t.X == 0, selection.Ternary(t.Y > 0, engine.PRESSED, engine.RELEASED), selection.Ternary(t.X > 0, engine.PRESSED, engine.RELEASED)).(byte)
				args := []interface{}{t.X, t.Y}
				appEngine.ProcessAction(action_byte, press_state, args)
			}
		}

		duration_us := time.Since(appEngine.LastFrame).Microseconds()

		if duration_us >= 16667 {
			appEngine.LastFrame = time.Now()
			appEngine.FrameTime = float64(duration_us) / 1e6

			err = appEngine.RenderScene()

			if err != nil {
				log.Fatalf("Error rendering scene: %s\n", err)
			}

		}
	}

	appEngine.FreeFonts()
}
