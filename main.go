package main

import (
	"log"
	"time"

	"main/engine"
	"main/selection"
	"main/utility"

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

	wind_width, wind_height := utility.Int32(800), utility.Int32(600)
	window, err := sdl.CreateWindow("Sudoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, wind_width.ToInt32(), wind_height.ToInt32(), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Error creating window: %s\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Error creating renderer: %s\n", err)
	}
	defer renderer.Destroy()

	appEngine := &engine.Engine{
		Window:    window,
		Renderer:  renderer,
		FrameTime: 0.0,
		MousePos:  utility.Vec2{X: &wind_width, Y: &wind_height},
	}

	appEngine.Setup()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				x_scal, y_scal := utility.Int32(t.X), utility.Int32(t.Y)
				appEngine.MousePos = utility.Vec2{
					X: &x_scal,
					Y: &y_scal,
				}
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

			err = appEngine.Renderer.Clear()

			if err != nil {
				log.Fatalf("Error clearing renderer: %s\n", err)
			}

			err = appEngine.DrawQuad(
				[4]sdl.Vertex{
					{
						Position: sdl.FPoint{X: 200, Y: 100},
						Color:    sdl.Color{R: uint8(0xFF), G: uint8(0x00), B: uint8(0x00), A: uint8(0xFF)},
						TexCoord: sdl.FPoint{},
					},
					{
						Position: sdl.FPoint{X: 200, Y: 500},
						Color:    sdl.Color{R: uint8(0x00), G: uint8(0xFF), B: uint8(0x00), A: uint8(0xFF)},
						TexCoord: sdl.FPoint{},
					},
					{
						Position: sdl.FPoint{X: 600, Y: 100},
						Color:    sdl.Color{R: uint8(0x00), G: uint8(0x00), B: uint8(0xFF), A: uint8(0xFF)},
						TexCoord: sdl.FPoint{},
					},
					{
						Position: sdl.FPoint{X: 600, Y: 500},
						Color:    sdl.Color{R: uint8(0xFF), G: uint8(0xFF), B: uint8(0xFF), A: uint8(0xFF)},
						TexCoord: sdl.FPoint{},
					},
				},
				[6]int32{0, 1, 2, 1, 2, 3},
			)

			if err != nil {
				log.Fatalf("Error drawing: %s\n", err)
			}

			font_name := "lotuscoder_normal"
			font_size := 24

			if err != nil {
				log.Fatalf("Error drawing text: %s\n", err)
			}

			appEngine.DrawText(
				font_name, font_size,
				[]string{
					"Hello, Traveler.",
					"Welcome to Sudoku.",
				},
				sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
				sdl.Point{X: 100, Y: 50},
			)

			appEngine.Renderer.Present()
		}
	}

	appEngine.FreeFonts()
}
